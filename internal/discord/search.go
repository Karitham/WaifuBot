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

type animeSearcher interface {
	Anime(string) ([]anilist.Media, error)
}

func (b *Bot) SearchAnime(w corde.ResponseWriter, i *corde.Interaction) {
	search, ok := i.Data.Options["search"].(string)
	if !ok {
		ephemeral(w, "please provide a search term")
	}

	anime, err := b.AnimeService.Anime(search)
	if err != nil {
		log.Err(err).Msg("error with anime service")
		ephemeral(w, "Error searching for this anime, either it doesn't exist or something went wrong")
		return
	}

	w.WithSource(&corde.InteractionRespData{
		Embeds: []corde.Embed{buildMediaEmbed(anime[0])},
	})
}

type mangaSearcher interface {
	Manga(string) ([]anilist.Media, error)
}

func (b *Bot) SearchManga(w corde.ResponseWriter, i *corde.Interaction) {
	search, ok := i.Data.Options["search"].(string)
	if !ok {
		ephemeral(w, "please provide a search term")
		return
	}

	manga, err := b.AnimeService.Manga(search)
	if err != nil || len(manga) < 1 {
		log.Err(err).Msg("error with anime service")
		ephemeral(w, "Error searching for this manga, either it doesn't exist or something went wrong")
		return
	}

	w.WithSource(&corde.InteractionRespData{
		Embeds: []corde.Embed{buildMediaEmbed(manga[0])},
	})
}

type userSearcher interface {
	User(string) ([]anilist.User, error)
}

func (b *Bot) SearchUser(w corde.ResponseWriter, i *corde.Interaction) {
	search, ok := i.Data.Options["search"].(string)
	if !ok {
		ephemeral(w, "please provide a search term")
		return
	}

	user, err := b.AnimeService.User(search)
	if err != nil && len(user) < 1 {
		log.Err(err).Msg("error with anime service")
		ephemeral(w, "Error searching for this user, either it doesn't exist or something went wrong")
		return
	}

	w.WithSource(&corde.InteractionRespData{
		Embeds: []corde.Embed{buildUserEmbed(user[0])},
	})
}

type charSearcher interface {
	Character(string) ([]anilist.Character, error)
}

func (b *Bot) SearchChar(w corde.ResponseWriter, i *corde.Interaction) {
	search, ok := i.Data.Options["search"].(string)
	if !ok {
		ephemeral(w, "please provide a search term")
		return
	}

	char, err := b.AnimeService.Character(search)
	if err != nil || len(char) < 1 {
		log.Err(err).Msg("error with anime service")
		ephemeral(w, "Error searching for this character, either it doesn't exist or something went wrong")
		return
	}

	w.WithSource(&corde.InteractionRespData{
		Embeds: []corde.Embed{buildCharEmbed(char[0])},
	})
}

func buildMediaEmbed(m anilist.Media) corde.Embed {
	return corde.Embed{
		Title:       FixString(m.Title.Romaji),
		Description: Sanitize(m.Description),
		URL:         m.Siteurl,
		Thumbnail:   corde.Image{URL: m.CoverImage.Large},
		Color:       ColorToInt(m.CoverImage.Color),
		Footer: corde.Footer{
			Text:    "View on anilist",
			IconURL: anilist.IconURL,
		},
		Image: corde.Image{URL: m.BannerImage},
	}
}

func buildUserEmbed(u anilist.User) corde.Embed {
	return corde.Embed{
		Title:       u.Name,
		Description: Sanitize(u.About),
		URL:         u.Siteurl,
		Thumbnail:   corde.Image{URL: u.Avatar.Large},
		// Anilist blue
		Color: anilist.Color,
		Footer: corde.Footer{
			Text:    "View on anilist",
			IconURL: anilist.IconURL,
		},
		Image: corde.Image{
			URL: fmt.Sprintf("https://img.anili.st/user/%d", u.ID),
		},
	}
}

func buildCharEmbed(c anilist.Character) corde.Embed {
	return corde.Embed{
		Title:       FixString(c.Name.Full),
		Description: Sanitize(c.Description),
		URL:         c.SiteURL,
		Thumbnail:   corde.Image{URL: c.Image.Large},
		// Anilist blue
		Color: anilist.Color,
		Footer: corde.Footer{
			Text:    "View on anilist",
			IconURL: anilist.IconURL,
		},
	}
}

// SanitizeHTML removes all HTML tags from the given string.
// It also removes double newlines and double || characters.
var SanitizeHTML = regexp.MustCompile(`<[^>]*>|\|\|[^|]*\|\||\s{2,}|img[\d\%]*\([^)]*\)|[#~*]{2,}|\n`)

// Sanitize removes all HTML tags from the given string.
// It also removes double newlines and double || characters.
func Sanitize(s string) string {
	return SanitizeHTML.ReplaceAllString(s, " ")
}

// FixString removes eventual
// double space or any whitespace possibly in a string
// and replace it with a space.
func FixString(s string) string {
	return strings.Join(strings.Fields(s), " ")
}

// ColorToInt
// Turn an hex color string beginning with a # into a uint32 representing a color.
func ColorToInt(s string) int64 {
	s = strings.Trim(s, "#")
	u, _ := strconv.ParseInt(s, 16, 64)
	return u
}
