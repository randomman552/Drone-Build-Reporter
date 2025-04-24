package types

type NotifyMode string

const (
	Finished   NotifyMode = "finished"
	InProgress NotifyMode = "started"
)

type Config struct {
	// Directory to look for go templates
	TemplateDirectory string
	// App token for gotify messages
	GotifyToken string
	// URL of gotify installation, should include scheme https://
	GotifyUrl string
	// Discord webhook to call
	DiscordWebhook string
	// The mode of notification
	NotifyMode NotifyMode
}
