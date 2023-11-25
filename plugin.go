package main

import (
	"reporter/reporters"
	"reporter/types"
)

type Plugin struct {
	Context types.DroneContext // Context related to the drone pipeline we are running in
	Config  types.Config       // Config for this plugin
}

func (p Plugin) Run() {
	reporters := []types.Reporter{
		reporters.ConsoleReporter{Config: p.Config},
		reporters.GotifyReporter{Config: p.Config},
	}

	for _, reporter := range reporters {
		reporter.Report(p.Context)
	}
}
