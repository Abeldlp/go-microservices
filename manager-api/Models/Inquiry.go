package models

import "gorm.io/gorm"

type Inquiry struct {
	gorm.Model
	Message  string `json:"message"`
	Deadline string `json:"deadline"`
	Assignee string `json:"assignee"`
}
