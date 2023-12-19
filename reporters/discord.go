package reporters

import (
	"log"
	"reporter/types"
)

type DiscordReporter struct {
	Config types.Config
}

func (r DiscordReporter) Report(context types.DroneContext) {
	if len((r.Config.DiscordWebhook)) <= 0 {
		log.Println(("Missing Discord Webhook"))
		return
	}
}
