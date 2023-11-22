package main

import (
	reporters "build-reporter/reporters"
)

func main() {
	context := reporters.DroneContext{}

	reporter := new(reporters.ConsoleReporter)
	reporter.Report(context)
}
