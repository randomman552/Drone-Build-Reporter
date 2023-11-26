package types

type Config struct {
	// Directory to look for go templates
	TemplateDirectory string
	// App token for gotify messages
	GotifyToken string
	// URL of gotify installation, should include scheme https://
	GotifyUrl string
}
