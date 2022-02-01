package anilist

import (
	"context"
	"errors"
	"fmt"
	"math/rand"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/Karitham/WaifuBot/internal/discord"
	"github.com/Khan/genqlient/graphql"
	"github.com/rs/zerolog/log"
)

// github.com/Khan/genqlient
//go:generate genqlient genqlient.yaml

// Anilist defines a common interface to interact with anilist
type Anilist struct {
	c             graphql.Client
	seed          rand.Source64
	MaxChars      int64
	internalCache map[string]querier
}

// Check that anilist actually implements the interface
var _ discord.TrackingService = (*Anilist)(nil)

type querier struct {
	*sync.Mutex
	cache map[any]any
}

// New returns a new anilist client
func New() *Anilist {
	const graphURL = "https://graphql.anilist.co"

	a := &Anilist{
		c:        graphql.NewClient(graphURL, &http.Client{Timeout: 5 * time.Second}),
		MaxChars: 50_000,
		seed:     rand.New(rand.NewSource(time.Now().Unix())),
		internalCache: map[string]querier{
			"random": {
				cache: make(map[any]any),
				Mutex: &sync.Mutex{},
			},
		},
	}

	go a.randomCache(a.internalCache["random"])
	return a
}

func (a *Anilist) RandomChar(ctx context.Context, notIn ...int64) (discord.MediaCharacter, error) {
	c := a.internalCache["random"]
	c.Lock()
	defer c.Unlock()

	rest := []discord.MediaCharacter{}

two:
	for id, char := range c.cache {
		for _, n := range notIn {
			if n == id {
				continue two
			}
		}

		rest = append(rest, char.(discord.MediaCharacter))
	}

	if len(c.cache) < 100 {
		for i := 0; i < 5; i++ {
			go func() {
				ch, err := a.randomChar(context.Background(), notIn...)
				if err != nil {
					log.Warn().Err(err).Msg("error getting random char")
					return
				}
				c.Lock()
				defer c.Unlock()
				c.cache[ch.ID] = ch
			}()
		}
	}

	if len(rest) > 0 {
		char := rest[a.seed.Int63()%(int64(len(rest)))]
		log.Trace().Str("char", char.Name).Int("cache size", len(c.cache)).Msg("Hit cache")
		delete(c.cache, char.ID)
		return char, nil
	}

	return a.randomChar(ctx, notIn...)
}

func (a *Anilist) randomChar(ctx context.Context, notIn ...int64) (discord.MediaCharacter, error) {
	r, err := charactersRandom(ctx, a.c, a.seed.Int63()%a.MaxChars, notIn)
	if err != nil {
		return discord.MediaCharacter{}, err
	}

	if len(r.Page.Characters) < 1 {
		return discord.MediaCharacter{}, errors.New("error querying random char")
	}

	c := r.Page.Characters[0]
	mediaTitle := ""
	if len(c.Media.Nodes) > 0 {
		mediaTitle = c.Media.Nodes[0].Title.Romaji
	}

	return discord.MediaCharacter{
		ID:         c.Id,
		Name:       strings.Join(strings.Fields(c.Name.Full), " "),
		ImageURL:   c.Image.Large,
		URL:        c.SiteUrl,
		MediaTitle: strings.Join(strings.Fields(mediaTitle), " "),
	}, nil
}

func (a *Anilist) Anime(ctx context.Context, title string) ([]discord.Media, error) {
	return a.media(ctx, title, MediaTypeAnime)
}

func (a *Anilist) media(ctx context.Context, title string, t MediaType) ([]discord.Media, error) {
	media, err := media(ctx, a.c, title, MediaTypeAnime)
	if err != nil {
		return nil, err
	}
	resp := make([]discord.Media, len(media.Page.Media))
	for i, m := range media.Page.Media {
		resp[i] = discord.Media{
			CoverImageURL:   m.CoverImage.Large,
			BannerImageURL:  m.BannerImage,
			CoverImageColor: ColorUint(m.CoverImage.Color),
			Title:           m.Title.Romaji,
			URL:             m.SiteUrl,
			Description:     m.Description,
		}
	}

	return resp, err
}

func (a *Anilist) User(ctx context.Context, name string) ([]discord.TrackerUser, error) {
	users, err := user(ctx, a.c, name)
	if err != nil {
		return nil, err
	}

	resp := make([]discord.TrackerUser, len(users.Page.Users))
	for i, u := range users.Page.Users {
		resp[i] = discord.TrackerUser{
			URL:      u.SiteUrl,
			Name:     u.Name,
			ImageURL: fmt.Sprintf("https://img.anili.st/user/%d", u.Id),
			About:    u.About,
		}
	}

	return resp, nil
}

func (a *Anilist) Character(ctx context.Context, name string) ([]discord.MediaCharacter, error) {
	char, err := character(ctx, a.c, name)
	if err != nil {
		return nil, err
	}

	resp := make([]discord.MediaCharacter, len(char.Page.Characters))
	for i, c := range char.Page.Characters {
		resp[i] = discord.MediaCharacter{
			ID:          c.Id,
			Name:        c.Name.Full,
			ImageURL:    c.Image.Large,
			URL:         c.SiteUrl,
			Description: c.Description,
		}
	}

	return resp, nil
}

func (a *Anilist) Manga(ctx context.Context, title string) ([]discord.Media, error) {
	return a.media(ctx, title, MediaTypeManga)
}

// ColorUint
// Turn an hex color string beginning with a # into a uint32 representing a color.
func ColorUint(s string) uint32 {
	s = strings.Trim(s, "#")
	u, _ := strconv.ParseUint(s, 16, 32)
	return uint32(u)
}

// SanitizeHTML removes all HTML tags from the given string.
// It also removes double newlines and double || characters.
var SanitizeHTML = regexp.MustCompile(`<[^>]*>|\|\|[^|]*\|\||\s{2,}|img[\d\%]*\([^)]*\)|[#~*]{2,}|\n`)

// Sanitize removes all HTML tags from the given string.
// It also removes double newlines and double || characters.
func Sanitize(s string) string {
	return SanitizeHTML.ReplaceAllString(s, "")
}

// FixString removes eventual
// double space or any whitespace possibly in a string
// and replace it with a space.
func FixString(s string) string {
	return strings.Join(strings.Fields(s), " ")
}

func (a *Anilist) randomCache(c querier) {
	for i := 0; i < 5; i++ {
		time.Sleep(500 * time.Millisecond)
		go func() {
			ch, err := a.randomChar(context.Background())
			if err != nil {
				return
			}
			c.Lock()
			defer c.Unlock()
			c.cache[ch.ID] = ch
		}()
	}
}
