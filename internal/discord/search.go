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
	search := i.Data.Options.String("search")

	anime, err := b.AnimeService.Anime(search)
	if err != nil {
		log.Err(err).Msg("error with anime service")
		w.Respond(corde.NewResp().Content("Error searching for this anime, either it doesn't exist or something went wrong").Ephemeral().B())
		return
	}

	w.Respond(corde.NewResp().Embeds(mediaEmbed(anime[0])).B())
}

type mangaSearcher interface {
	Manga(string) ([]anilist.Media, error)
}

func (b *Bot) SearchManga(w corde.ResponseWriter, i *corde.Interaction) {
	search := i.Data.Options.String("search")

	manga, err := b.AnimeService.Manga(search)
	if err != nil || len(manga) < 1 {
		log.Err(err).Msg("error with manga service")
		w.Respond(corde.NewResp().Content("Error searching for this manga, either it doesn't exist or something went wrong").Ephemeral().B())
		return
	}

	w.Respond(corde.NewResp().Embeds(mediaEmbed(manga[0])).B())
}

type userSearcher interface {
	User(string) ([]anilist.User, error)
}

func (b *Bot) SearchUser(w corde.ResponseWriter, i *corde.Interaction) {
	search := i.Data.Options.String("search")

	user, err := b.AnimeService.User(search)
	if err != nil || len(user) < 1 {
		log.Err(err).Msg("error with user service")
		w.Respond(corde.NewResp().Content("Error searching for this user, either it doesn't exist or something went wrong").Ephemeral().B())
		return
	}

	w.Respond(corde.NewResp().Embeds(userEmbed(user[0])).B())
}

type charSearcher interface {
	Character(string) ([]anilist.Character, error)
}

func (b *Bot) SearchChar(w corde.ResponseWriter, i *corde.Interaction) {
	search := i.Data.Options.String("search")

	char, err := b.AnimeService.Character(search)
	if err != nil || len(char) < 1 {
		log.Err(err).Msg("error with char service")
		w.Respond(corde.NewResp().Content("Error searching for this character, either it doesn't exist or something went wrong").Ephemeral().B())
		return
	}

	w.Respond(corde.NewResp().Embeds(charEmbed(char[0])).B())
}

func mediaEmbed(m anilist.Media) corde.Embed {
	return applyEmbedOpt(corde.NewEmbed().
		Title(FixString(m.Title.Romaji)).
		URL(m.Siteurl).
		Color(ColorToInt(m.CoverImage.Color)).
		Image(corde.Image{URL: m.BannerImage}).
		Thumbnail(corde.Image{URL: m.CoverImage.Large}),

		description(m.Description),
		anilistFooter,
	).B()
}

func userEmbed(u anilist.User) corde.Embed {
	return applyEmbedOpt(corde.NewEmbed().
		Title(FixString(u.Name)).
		URL(u.Siteurl).
		Color(anilist.Color).
		Image(corde.Image{URL: fmt.Sprintf("https://img.anili.st/user/%d", u.ID)}),

		description(u.About),
		anilistFooter,
	).B()
}

func charEmbed(c anilist.Character) corde.Embed {
	return applyEmbedOpt(corde.NewEmbed().
		Title(FixString(c.Name.Full)).
		Color(anilist.Color).
		URL(c.SiteURL).
		Thumbnail(corde.Image{URL: c.Image.Large}),

		description(c.Description),
		anilistFooter,
	).B()
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
