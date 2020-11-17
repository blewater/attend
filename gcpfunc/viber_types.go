package gcpfunc

type MessageReply struct {
	Receiver     string    `json:"receiver,omitempty"`
	Sender       *User     `json:"sender,omitempty"`
	TrackingData string    `json:"tracking_data,omitempty"`
	Type         string    `json:"type,omitempty"`
	Text         string    `json:"text,omitempty"`
	Media        string    `json:"media,omitempty"`
	Thumbnail    string    `json:"thumbnail,omitempty"`
	Keyboard     *Keyboard `json:"keyboard,omitempty"`
}

type ViberRequest struct {
	Event        string           `json:"event,omitempty"`
	Timestamp    int64            `json:"timestamp,omitempty"`
	ChatHostname string           `json:"chat_hostname,omitempty"`
	MessageToken int64            `json:"message_token,omitempty"`
	Type         string           `json:"type,omitempty"`
	User         *User            `json:"person,omitempty"`
	Sender       *User            `json:"sender,omitempty"`
	Subscribed   bool             `json:"subscribed,omitempty"`
	Message      *MessageReceived `json:"message,omitempty"`
}

type User struct {
	ID         string `json:"id,omitempty"`
	Name       string `json:"name,omitempty"`
	Avatar     string `json:"avatar,omitempty"`
	Language   string `json:"language,omitempty"`
	Country    string `json:"country,omitempty"`
	APIVersion int    `json:"api_version,omitempty"`
}

type Location struct {
	Lat float64 `json:"lat,omitempty"`
	Lon float64 `json:"lon,omitempty"`
}

type MessageReceived struct {
	Type         string    `json:"type,omitempty"`
	Text         string    `json:"text,omitempty"`
	Media        string    `json:"media,omitempty"`
	Location     *Location `json:"location,omitempty"`
	TrackingData string    `json:"tracking_data,omitempty"`
}

type RegResp struct {
	Status        int      `json:"status"`
	StatusMessage string   `json:"status_message"`
	EventTypes    []string `json:"event_types"`
}

type Keyboard struct {
	Type          string   `json:"Type"`
	DefaultHeight bool     `json:"DefaultHeight"`
	Buttons       []Button `json:"Buttons"`
}

type Button struct {
	ActionType string `json:"ActionType"`
	ActionBody string `json:"ActionBody"`
	Text       string `json:"Text"`
	TextSize   string `json:"TextSize"`
	Columns    int    `json:"Columns"`
	Rows       int    `json:"Rows"`
	TextHAlign string `json:"TextHAlign"`
	TextVAlign string `json:"TextVAlign"`
	BgColor    string `json:"BgColor"`
	Image      string `json:"Image"`
	Silent     bool   `json:"Silent"`
}
