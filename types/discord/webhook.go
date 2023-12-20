package discord

type Webhook struct {
	Content    string        `json:"content"`
	Components []interface{} `json:"components"`
}

// Create a new Webhook message with the given content
func NewWebhook(content string) *Webhook {
	return &Webhook{
		Content:    content,
		Components: []interface{}{},
	}
}

// Add a component to the webhook components
func (w *Webhook) AppendComponent(component interface{}) {
	w.Components = append(w.Components, component)
}
