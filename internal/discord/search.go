package discord

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/Karitham/WaifuBot/internal/anilist"
	"github.com/Karitham/corde"
	"github.com/rs/zerolog/log"
)

func (b *Bot) search(m *corde.Mux) {
	m.Command("char", trace(b.SearchChar))
	m.Command("user", trace(b.SearchUser))
	m.Command("manga", trace(b.SearchManga))
	m.Command("anime", trace(b.SearchAnime))
}

type animeSearcher interface {
	Anime(string) ([]anilist.Media, error)
}

func (b *Bot) SearchAnime(w corde.ResponseWriter, i *corde.InteractionRequest) {
	search, _ := i.Data.Options.String("name")

	anime, err := b.AnimeService.Anime(search)
	if err != nil {
		log.Err(err).Msg("error with anime service")
		w.Respond(rspErr("Error searching for this anime, either it doesn't exist or something went wrong"))
		return
	}

	w.Respond(mediaEmbed(anime[0]))
}

type mangaSearcher interface {
	Manga(string) ([]anilist.Media, error)
}

func (b *Bot) SearchManga(w corde.ResponseWriter, i *corde.InteractionRequest) {
	search, _ := i.Data.Options.String("name")

	manga, err := b.AnimeService.Manga(search)
	if err != nil || len(manga) < 1 {
		log.Err(err).Msg("error with manga service")
		w.Respond(rspErr("Error searching for this manga, either it doesn't exist or something went wrong"))
		return
	}

	w.Respond(mediaEmbed(manga[0]))
}

type userSearcher interface {
	User(string) ([]anilist.User, error)
}

func (b *Bot) SearchUser(w corde.ResponseWriter, i *corde.InteractionRequest) {
	search, _ := i.Data.Options.String("name")

	user, err := b.AnimeService.User(search)
	if err != nil || len(user) < 1 {
		log.Err(err).Msg("error with user service")
		w.Respond(rspErr("Error searching for this user, either it doesn't exist or something went wrong"))
		return
	}

	w.Respond(userEmbed(user[0]))
}

type charSearcher interface {
	Character(string) ([]anilist.Character, error)
}

func (b *Bot) SearchChar(w corde.ResponseWriter, i *corde.InteractionRequest) {
	search, _ := i.Data.Options.String("name")

	char, err := b.AnimeService.Character(search)
	if err != nil || len(char) < 1 {
		log.Err(err).Msg("error with char service")
		w.Respond(rspErr("Error searching for this character, either it doesn't exist or something went wrong"))
		return
	}

	w.Respond(charEmbed(char[0]))
}

func mediaEmbed(m anilist.Media) *corde.EmbedB {
	return applyEmbedOpt(corde.NewEmbed().
		Title(FixString(m.Title.Romaji)).
		URL(m.Siteurl).
		Color(ColorToInt(m.CoverImage.Color)).
		Image(corde.Image{URL: m.BannerImage}).
		Thumbnail(corde.Image{URL: m.CoverImage.Large}),

		description(m.Description),
		anilistFooter,
	)
}

func userEmbed(u anilist.User) *corde.EmbedB {
	return applyEmbedOpt(corde.NewEmbed().
		Title(FixString(u.Name)).
		URL(u.Siteurl).
		Color(anilist.Color).
		Image(corde.Image{URL: fmt.Sprintf("https://img.anili.st/user/%d", u.ID)}),

		description(u.About),
		anilistFooter,
	)
}

func charEmbed(c anilist.Character) *corde.EmbedB {
	return applyEmbedOpt(corde.NewEmbed().
		Title(FixString(c.Name.Full)).
		Color(anilist.Color).
		URL(c.SiteURL).
		Thumbnail(corde.Image{URL: c.Image.Large}),

		description(c.Description),
		anilistFooter,
	)
}

func anilistFooter(b *corde.EmbedB) *corde.EmbedB {
	return b.Footer(corde.Footer{
		Text:    "View on anilist",
		IconURL: anilist.IconURL,
	})
}

func description(d string) func(*corde.EmbedB) *corde.EmbedB {
	return func(b *corde.EmbedB) *corde.EmbedB { return b.Description(Sanitize(d)) }
}

func applyEmbedOpt(b *corde.EmbedB, opts ...func(*corde.EmbedB) *corde.EmbedB) *corde.EmbedB {
	for _, opt := range opts {
		b = opt(b)
	}
	return b
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

// ColorToInt
// Turn an hex color string beginning with a # into a uint32 representing a color.
func ColorToInt(s string) uint32 {
	s = strings.Trim(s, "#")
	u, _ := strconv.ParseUint(s, 16, 32)
	return uint32(u)
}
