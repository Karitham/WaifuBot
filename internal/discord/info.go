package discord

import (
	"runtime"
	"strconv"
	"strings"
	"time"

	"github.com/Karitham/corde"
	"github.com/Karitham/corde/components"
)

func (b *Bot) info(w corde.ResponseWriter, i *corde.Request[components.SlashCommandInteractionData]) {
	infob := strings.Builder{}
	infob.WriteString("Version: ")
	infob.WriteString(runtime.Version())
	infob.WriteRune('\n')

	infob.WriteString("Runtime: ")
	infob.WriteString(runtime.GOOS + "/" + runtime.GOARCH)
	infob.WriteRune('\n')

	infob.WriteString("Goroutines: ")
	infob.WriteString(strconv.Itoa(runtime.NumGoroutine()))
	infob.WriteRune('\n')

	s := &runtime.MemStats{}
	runtime.ReadMemStats(s)
	infob.WriteString("Memory: ")
	infob.WriteString(strconv.Itoa(int(s.Alloc / 1024)))
	infob.WriteString(" KB")
	infob.WriteRune('\n')

	infob.WriteString("GC Pauses: ")
	infob.WriteString(time.Duration(int64(s.PauseTotalNs)).String())
	infob.WriteRune('\n')

	w.Respond(components.NewResp().Ephemeral().Embeds(
		components.NewEmbed().
			Title("Info").
			Descriptionf(
				"A gacha game bot to collect and trade characters, and discover anything manga related.\n```info\n%s```",
				infob.String(),
			).
			URL("https://github.com/Karitham/WaifuBot/tree/corde"),
	))
}
