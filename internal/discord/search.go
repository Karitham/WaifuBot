package discord

import (
	"context"
	"regexp"

	"github.com/Karitham/corde"
	"github.com/rs/zerolog/log"
)

func (b *Bot) search(m *corde.Mux) {
	t := trace[corde.SlashCommandInteractionData]
	i := interact(b.Inter, onInteraction[corde.SlashCommandInteractionData](b))

	m.SlashCommand("char", wrap(b.SearchChar, t, i))
	m.SlashCommand("user", wrap(b.SearchUser, t, i))
	m.SlashCommand("manga", wrap(b.SearchManga, t, i))
	m.SlashCommand("anime", wrap(b.SearchAnime, t, i))
}

type AnimeSearcher interface {
	Anime(context.Context, string) ([]Media, error)
}

// Media represents an anime or manga.
type Media struct {
	Title           string
	URL             string
	CoverImageURL   string
	BannerImageURL  string
	CoverImageColor uint32
	Description     string
}

// TrackerUser represents an anime tracker user.
type TrackerUser struct {
	Name     string
	URL      string
	ImageURL string
	About    string
}

func (b *Bot) SearchAnime(w corde.ResponseWriter, i *corde.Request[corde.SlashCommandInteractionData]) {
	search, _ := i.Data.Options.String("name")

	anime, err := b.AnimeService.Anime(i.Context, search)
	if err != nil {
		log.Ctx(i.Context).Err(err).Msg("error with anime service")
		w.Respond(rspErr("Error searching for this anime, either it doesn't exist or something went wrong"))
		return
	}

	w.Respond(mediaEmbed(anime[0]))
}

type MangaSearcher interface {
	Manga(context.Context, string) ([]Media, error)
}

func (b *Bot) SearchManga(w corde.ResponseWriter, i *corde.Request[corde.SlashCommandInteractionData]) {
	search, _ := i.Data.Options.String("name")

	manga, err := b.AnimeService.Manga(i.Context, search)
	if err != nil || len(manga) < 1 {
		log.Ctx(i.Context).Err(err).Msg("error with manga service")
		w.Respond(rspErr("Error searching for this manga, either it doesn't exist or something went wrong"))
		return
	}

	w.Respond(mediaEmbed(manga[0]))
}

type UserSearcher interface {
	User(context.Context, string) ([]TrackerUser, error)
}

func (b *Bot) SearchUser(w corde.ResponseWriter, i *corde.Request[corde.SlashCommandInteractionData]) {
	search, _ := i.Data.Options.String("name")

	user, err := b.AnimeService.User(i.Context, search)
	if err != nil || len(user) < 1 {
		log.Ctx(i.Context).Err(err).Msg("error with user service")
		w.Respond(rspErr("Error searching for this user, either it doesn't exist or something went wrong"))
		return
	}

	w.Respond(userEmbed(user[0]))
}

type CharSearcher interface {
	Character(context.Context, string) ([]MediaCharacter, error)
}

func (b *Bot) SearchChar(w corde.ResponseWriter, i *corde.Request[corde.SlashCommandInteractionData]) {
	search, _ := i.Data.Options.String("name")

	char, err := b.AnimeService.Character(i.Context, search)
	if err != nil || len(char) < 1 {
		log.Ctx(i.Context).Err(err).Msg("error with char service")
		w.Respond(rspErr("Error searching for this character, either it doesn't exist or something went wrong"))
		return
	}

	w.Respond(charEmbed(char[0]))
}

func mediaEmbed(m Media) *corde.EmbedB {
	return applyEmbedOpt(corde.NewEmbed().
		Title(m.Title).
		URL(m.URL).
		Color(m.CoverImageColor).
		ImageURL(m.BannerImageURL).
		Thumbnail(corde.Image{URL: m.CoverImageURL}).
		Description(m.Description),
		anilistFooter,
		sanitizeDescOpt,
	)
}

func userEmbed(u TrackerUser) *corde.EmbedB {
	return applyEmbedOpt(corde.NewEmbed().
		Title(u.Name).
		URL(u.URL).
		Color(AnilistColor).
		ImageURL(u.ImageURL).
		Description(u.About),
		anilistFooter,
		sanitizeDescOpt,
	)
}

func charEmbed(c MediaCharacter) *corde.EmbedB {
	return applyEmbedOpt(corde.NewEmbed().
		Title(c.Name).
		Color(AnilistColor).
		URL(c.URL).
		Thumbnail(corde.Image{URL: c.ImageURL}).
		Description(c.Description),
		anilistFooter,
		sanitizeDescOpt,
	)
}

func anilistFooter(b *corde.EmbedB) *corde.EmbedB {
	return b.Footer(corde.Footer{
		Text:    "View on anilist",
		IconURL: AnilistIconURL,
	})
}

func applyEmbedOpt(b *corde.EmbedB, opts ...func(*corde.EmbedB) *corde.EmbedB) *corde.EmbedB {
	for _, opt := range opts {
		b = opt(b)
	}
	return b
}

func sanitizeDescOpt(b *corde.EmbedB) *corde.EmbedB {
	return b.Description(sanitizeDesc(b.Embed().Description))
}

var descRegexp = regexp.MustCompile(`!~|~!`)

func sanitizeDesc(desc string) string {
	return descRegexp.ReplaceAllString(desc, "||")
}
