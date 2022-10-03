package models

type Mail struct {
	Subject string `json:"subject"`
	Message string `json:"message"`
	From    string `json:"from"`
}
