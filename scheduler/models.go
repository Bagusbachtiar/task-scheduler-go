package scheduler

import "time"

// ScheduleTask matches our Neon Cloud Database Schema layout
type ScheduleTask struct {
	ID          int
	PhoneNumber string
	Message     string
	ExecutionTZ time.Time
}

// Language structures Meta's language code payload
type Language struct {
	Code string `json:"code"`
}

// TemplateDetails maps the inner template fields safely
type TemplateDetails struct {
	Name     string   `json:"name"`
	Language Language `json:"language"`
}

// WhatsAppTemplateRequest structures the exact JSON format Meta's API expects
type WhatsAppTemplateRequest struct {
	MessagingProduct string          `json:"messaging_product"`
	To               string          `json:"to"`
	Type             string          `json:"type"`
	Template         TemplateDetails `json:"template"`
}
