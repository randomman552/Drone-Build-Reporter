package types

type MessageRequestExtrasNotificationClick struct {
	Url string `json:"url"`
}

type MessageRequestExtrasNotification struct {
	Click MessageRequestExtrasNotificationClick `json:"click"`
}

type MessageRequestExtrasDisplay struct {
	ContentType string `json:"contentType"`
}

type MessageRequestExtras struct {
	Display      MessageRequestExtrasDisplay      `json:"client::display"`
	Notification MessageRequestExtrasNotification `json:"client::notification"`
}

type MessageRequest struct {
	Title    string               `json:"title"`
	Message  string               `json:"message"`
	Priority int                  `json:"priority"`
	Extras   MessageRequestExtras `json:"extras"`
}
