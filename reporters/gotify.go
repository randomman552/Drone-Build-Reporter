package reporters

import (
	"bytes"
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
	tplate, err := template.ParseFiles("templates/gotify/request.tmpl", "templates/gotify/message.tmpl")

	if err != nil {
		panic(err)
	}

	// Render message template
	messageBuffer := &bytes.Buffer{}
	tplate.ExecuteTemplate(messageBuffer, "message", context)
	messageBytes, err := io.ReadAll(messageBuffer)
	messageString := string(messageBytes)

	// Render request template
	gotifyContext := types.GotifyRequestContext{
		Drone:   context,
		Message: messageString,
	}
	buffer := &bytes.Buffer{}
	tplate.ExecuteTemplate(buffer, "request", gotifyContext)

	return buffer
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
