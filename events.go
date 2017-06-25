package slack

const MessageType = "message"
const URLVerificationType = "url_verification"

type EventCallback struct {
	Token       string   `json:"token"`
	TeamID      string   `json:"team_id"`
	ApiAppID    string   `json:"api_app_id"`
	Type        string   `json:"type"`
	AuthedUsers []string `json:"authed_users"`
	EventID     string   `json:"event_id"`
	EventTime   int64    `json:"event_time"`
	Challenge   *string  `json:"challenge"`
}

type EventMsg struct {
	EventCallback
	Event Msg `json:"event"`
}
