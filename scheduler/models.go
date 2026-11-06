package scheduler

import "time"

//ScheduleTask matches our neon Cloud Database Schema Layout
type ScheduleTask struct {
	ID 		 int
	PhoneNumber string
	Message     string
	ExecutionTZ time.Time //the exact time this task should fire
}

//WhatsappTextRequest structures the JSON payload for the Whatsapp Cloud API
type WhatsappTextRequest struct {
	MessagingProduct string `json:"messaging_product"`
	To               string `json:"to"`
	Type             string `json:"type"`
	Text             struct {
		Body string `json:"body"`
	} `json:"text"`
}