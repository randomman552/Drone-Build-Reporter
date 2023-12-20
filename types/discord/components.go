package discord

// Base discord message component
type Component struct {
	Type uint8 `json:"type"`
}

// Struct representing an ActionRow message component
type ActionRowComponent struct {
	Type       uint8         `json:"type"`
	Components []interface{} `json:"components"`
}

// Create a new ActionRowComponent
func NewActionRow() *ActionRowComponent {
	return &ActionRowComponent{
		Type:       1,
		Components: []interface{}{},
	}
}

// Add a component to the webhook components
func (c *ActionRowComponent) AppendComponent(component interface{}) {
	c.Components = append(c.Components, component)
}

// Struct representing a Button message component
type ButtonComponent struct {
	Type  uint8  `json:"type"`
	Label string `json:"label"`
	Url   string `json:"url"`
}

// Create a new ButtonComponent
func NewButton(label string, url string) *ButtonComponent {
	return &ButtonComponent{
		Type:  2,
		Label: label,
		Url:   url,
	}
}
