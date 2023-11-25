package reporters

import (
	"fmt"
	"reporter/types"
)

type GotifyReporter struct {
	Config types.Config
}

func (r GotifyReporter) Report(context types.DroneContext) {
	fmt.Println("Using URL:", r.Config.GotifyUrl)
}
