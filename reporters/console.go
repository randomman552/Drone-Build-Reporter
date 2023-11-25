package reporters

import (
	"os"
	"reporter/types"
	"text/template"
)

type ConsoleReporter struct {
	Config types.Config
}

func (c ConsoleReporter) Report(context types.DroneContext) {
	tplate, err := template.ParseFiles("/templates/console.tmpl")

	if err != nil {
		panic(err)
	}

	err = tplate.Execute(os.Stdout, context)

	if err != nil {
		panic(err)
	}
}
