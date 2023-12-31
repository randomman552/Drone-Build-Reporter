package reporters

import (
	"log"
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
		log.Printf("Error reading console template: %s", err)
		return
	}

	err = tplate.Execute(os.Stdout, context)

	if err != nil {
		log.Printf("Error rendering console template: %s", err)
		return
	}
}
