package main

import (
	reporters "build-reporter/reporters"
)

func main() {
	context := reporters.ReporterContext{Message: "Hello World!"}

	reporter := new(reporters.ConsoleReporter)
	reporter.Report(context)
}
