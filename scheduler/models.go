package scheduler

import "time"

// ScheduleTask matches our Neon Cloud Database Schema layout
type ScheduleTask struct {
	ID          int
	PhoneNumber string
	Message     string
	ExecutionTZ time.Time
}

type Language struct {
	Code string `json:"code"`
}

// Parameter maps the internal template text parameters Meta requires
type Parameter struct {
	Type string `json:"type"`
	Text string `json:"text"`
}

// Component maps the parameter components array structure
type Component struct {
	Type       string      `json:"type"`
	Parameters []Parameter `json:"parameters"`
}

type TemplateDetails struct {
	Name       string      `json:"name"`
	Language   Language    `json:"language"`
	Components []Component `json:"components"` // ADDED: Matches Meta's expected array structure
}

type WhatsAppTemplateRequest struct {
	MessagingProduct string          `json:"messaging_product"`
	To               string          `json:"to"`
	Type             string          `json:"type"`
	Template         TemplateDetails `json:"template"`
}
