package reporters

import "fmt"

type ConsoleReporter struct{}

func (c ConsoleReporter) Report(context ReporterContext) {
	fmt.Println(context.Message)
}
