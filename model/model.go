package model

// Message defines model for Message.
type Message struct {
	Context     string         `json:"context"`
	CreatedAt   string         `json:"createdAt"`
	Description *string        `json:"description,omitempty"`
	Lang        string         `json:"lang"`
	Message     string         `json:"message"`
	MessageID   string         `json:"messageID"`
	Translation *[]Translation `json:"translation,omitempty"`
	UserID      string         `json:"userID"`
	Uuid        string         `json:"uuid"`
}

// MessagePayload defines model for MessagePayload.
type MessagePayload struct {
	Context     string  `json:"context"`
	Description *string `json:"description,omitempty"`
	Lang        string  `json:"lang"`
	Message     string  `json:"message"`
	MessageID   string  `json:"messageID"`
}

// Translation defines model for Translation.
type Translation struct {
	CreatedAt string `json:"createdAt"`
	Lang      string `json:"lang"`
	Message   string `json:"message"`
	UserID    string `json:"userID"`
	Uuid      string `json:"uuid"`
}
