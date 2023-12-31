package reporters

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"net/url"
	"path"
	"reporter/types"
	"text/template"
)

type GotifyReporter struct {
	Config types.Config
}

func (r GotifyReporter) GetUrl() *url.URL {
	url, err := url.Parse(r.Config.GotifyUrl)

	if err != nil {
		log.Fatalf("Error parsing gotify url: %s", err)
		return nil
	}

	url = url.JoinPath("/message")

	// Setup query string
	query := url.Query()
	query.Set("token", r.Config.GotifyToken)
	url.RawQuery = query.Encode()

	return url
}

func (r GotifyReporter) RenderTemplate(context types.DroneContext) *bytes.Buffer {
	templatePath := path.Join(r.Config.TemplateDirectory, "gotify.tmpl")
	tplate, err := template.ParseFiles(templatePath)

	if err != nil {
		log.Fatalf("Error parsing template: %s", err)
		return nil
	}

	// Build markdown from template
	messageBuffer := &bytes.Buffer{}
	tplate.Execute(messageBuffer, context)
	messageBytes, _ := io.ReadAll(messageBuffer)
	messageString := string(messageBytes)

	// Build title
	titleBuffer := &bytes.Buffer{}
	tplate.ExecuteTemplate(titleBuffer, "title", context)
	titleBytes, _ := io.ReadAll(titleBuffer)
	titleString := string(titleBytes)

	// Build JSON request
	request := types.MessageRequest{
		Title:    titleString,
		Message:  messageString,
		Priority: 5,
		Extras: types.MessageRequestExtras{
			Display: types.MessageRequestExtrasDisplay{
				ContentType: "text/markdown",
			},
			Notification: types.MessageRequestExtrasNotification{
				Click: types.MessageRequestExtrasNotificationClick{
					Url: context.Build.Link,
				},
			},
		},
	}
	requestBuffer := &bytes.Buffer{}
	json.NewEncoder(requestBuffer).Encode(request)

	return requestBuffer
}

func (r GotifyReporter) BuildRequest(context types.DroneContext) *http.Request {
	url := r.GetUrl()
	body := r.RenderTemplate(context)

	if url == nil || body == nil {
		return nil
	}

	request, err := http.NewRequest("POST", url.String(), body)

	if err != nil {
		log.Fatalf("Error building request: %s", err)
		return nil
	}

	request.Header.Set("Content-Type", "application/json")

	return request
}

func (r GotifyReporter) Report(context types.DroneContext) {
	if len(r.Config.GotifyUrl) <= 0 {
		log.Println("Missing Gotify URL")
		return
	}
	if len(r.Config.GotifyToken) <= 0 {
		log.Println("Missing Gotify token")
		return
	}

	request := r.BuildRequest(context)
	if request == nil {
		return
	}

	client := http.Client{}
	response, err := client.Do(request)

	if err != nil {
		log.Fatalf("Error making request: %s", err)
		return
	}

	if response.StatusCode != 200 {
		body, _ := io.ReadAll(response.Body)
		log.Fatal(response.Status, " - ", string(body))
	}

	response.Body.Close()
}
