package reporters

import (
	"os"
	"path"
	"reporter/types"
	"text/template"
)

type ConsoleReporter struct {
	Config types.Config
}

func (c ConsoleReporter) Report(context types.DroneContext) {
	templatePath := path.Join(c.Config.TemplateDirectory, "console.tmpl")
	tplate, err := template.ParseFiles(templatePath)

	if err != nil {
		panic(err)
	}

	err = tplate.Execute(os.Stdout, context)

	if err != nil {
		panic(err)
	}
}
