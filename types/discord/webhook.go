package discord

type Webhook struct {
	Embeds []Embed `json:"embeds"`
}

// Create a new Webhook message with the given content
func NewWebhook() *Webhook {
	return &Webhook{
		Embeds: []Embed{},
	}
}

// Add an Embed to the webhook Embeds
func (w *Webhook) AppendEmbed(embed Embed) {
	w.Embeds = append(w.Embeds, embed)
}

// Struct representing a discord Webhook embed
type Embed struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

// Create a new Discord Webhook Embed
func NewEmbed(title string, description string) *Embed {
	return &Embed{
		Title:       title,
		Description: description,
	}
}
