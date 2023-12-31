package reporters

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"path"
	"reporter/types"
	"reporter/types/discord"
	"text/template"
)

type DiscordReporter struct {
	Config types.Config
}

func (r DiscordReporter) RenderTemplate(context types.DroneContext) *bytes.Buffer {
	templatePath := path.Join(r.Config.TemplateDirectory, "discord.tmpl")
	tplate, err := template.ParseFiles(templatePath)
	buffer := &bytes.Buffer{}

	if err != nil {
		log.Printf("Error parsing discord template: %s", err)
		return nil
	}

	// Build markdown from template
	tplate.Execute(buffer, context)
	message := buffer.String()
	buffer.Reset()

	// Build title from template
	tplate.ExecuteTemplate(buffer, "title", context)
	title := buffer.String()
	buffer.Reset()

	// Build JSON request
	request := discord.NewWebhook()
	request.AppendEmbed(*discord.NewEmbed(title, message, context.Build.Link))

	// Package request body in a bytes buffer
	json.NewEncoder(buffer).Encode(request)
	return buffer
}

func (r DiscordReporter) BuildRequest(context types.DroneContext) *http.Request {
	url := r.Config.DiscordWebhook
	body := r.RenderTemplate(context)
	if body == nil {
		return nil
	}

	request, err := http.NewRequest("POST", url, body)

	if err != nil {
		log.Printf("Error building discord webhook request")
		return nil
	}

	request.Header.Set("Content-Type", "application/json")

	return request
}

func (r DiscordReporter) Report(context types.DroneContext) {
	if len((r.Config.DiscordWebhook)) <= 0 {
		log.Println(("Missing Discord Webhook... skipping..."))
		return
	}

	request := r.BuildRequest(context)

	client := http.Client{}
	response, err := client.Do(request)

	if err != nil {
		log.Printf("Error completing discord webhook request: %s", err)
		return
	}

	if response.StatusCode != 204 {
		body, _ := io.ReadAll(response.Body)
		log.Printf("Error completing discord webhook request: %s - %s", response.Status, string(body))
	}

	response.Body.Close()
}
