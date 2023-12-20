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

	if err != nil {
		panic(err)
	}

	// Build markdown from template
	messageBuffer := &bytes.Buffer{}
	tplate.Execute(messageBuffer, context)
	messageBytes, _ := io.ReadAll(messageBuffer)
	messageString := string(messageBytes)

	// Build links for message
	actionRow := discord.NewActionRow()
	actionRow.AppendComponent(discord.NewButton("View Commit", context.Commit.Link))
	actionRow.AppendComponent(discord.NewButton("View Build", context.Build.Link))
	actionRow.AppendComponent(discord.NewButton("View Repo", context.Repo.Link))

	// Build JSON request
	request := discord.NewWebhook(messageString)
	request.AppendComponent(actionRow)

	// Package request body in a bytes buffer
	requestBuffer := &bytes.Buffer{}
	json.NewEncoder(requestBuffer).Encode(request)

	return requestBuffer
}

func (r DiscordReporter) BuildRequest(context types.DroneContext) *http.Request {
	url := r.Config.DiscordWebhook
	body := r.RenderTemplate(context)
	request, err := http.NewRequest("POST", url, body)

	if err != nil {
		panic(err)
	}

	request.Header.Set("Content-Type", "application/json")

	return request
}

func (r DiscordReporter) Report(context types.DroneContext) {
	if len((r.Config.DiscordWebhook)) <= 0 {
		log.Println(("Missing Discord Webhook"))
		return
	}

	request := r.BuildRequest(context)

	client := http.Client{}
	response, err := client.Do(request)

	if err != nil {
		panic(err)
	}

	if response.StatusCode != 204 {
		body, _ := io.ReadAll(response.Body)
		log.Fatal(response.Status, " - ", string(body))
	}

	response.Body.Close()
}
