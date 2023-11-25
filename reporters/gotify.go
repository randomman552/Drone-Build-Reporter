package reporters

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"net/url"
	"reporter/types"
	"text/template"
)

type GotifyReporter struct {
	Config types.Config
}

func (r GotifyReporter) GetUrl() *url.URL {
	url, err := url.Parse(r.Config.GotifyUrl)

	if err != nil {
		panic(err)
	}

	url = url.JoinPath("/message")

	// Setup query string
	query := url.Query()
	query.Set("token", r.Config.GotifyToken)
	url.RawQuery = query.Encode()

	return url
}

func (r GotifyReporter) RenderTemplate(context types.DroneContext) *bytes.Buffer {
	tplate, err := template.ParseFiles("templates/gotify.tmpl")

	if err != nil {
		panic(err)
	}

	// Build markdown from template
	messageBuffer := &bytes.Buffer{}
	tplate.Execute(messageBuffer, context)
	messageBytes, _ := io.ReadAll(messageBuffer)
	messageString := string(messageBytes)

	// Build JSON request
	request := types.MessageRequest{
		Title:    context.Build.Status + " for " + context.Repo.Namespace + "/" + context.Repo.Name,
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
	request, err := http.NewRequest("POST", url.String(), body)

	if err != nil {
		panic(err)
	}

	request.Header.Set("Content-Type", "application/json")

	return request
}

func (r GotifyReporter) Report(context types.DroneContext) {
	if len(r.Config.GotifyUrl) <= 0 {
		log.Println("Missing URL")
		return
	}
	if len(r.Config.GotifyToken) <= 0 {
		log.Println("Missing token")
		return
	}

	request := r.BuildRequest(context)

	client := http.Client{}
	response, err := client.Do(request)

	if err != nil {
		panic(err)
	}

	if response.StatusCode != 200 {
		body, _ := io.ReadAll(response.Body)
		log.Fatal(response.Status, " - ", string(body))
	}

	response.Body.Close()
}
