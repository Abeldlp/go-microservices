package models

type Mail struct {
	Subject  string `json:"subject"`
	Sender   string `json:"sender"`
	Receiver string `json:"receiver"`
	Content  string `json:"content"`
}
