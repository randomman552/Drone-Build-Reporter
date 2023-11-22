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
	reporter := reporters.ConsoleReporter{Config: p.Config}

	reporter.Report(p.Context)
}
