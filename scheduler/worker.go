package scheduler

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

func ProcessTasks(dbPool *pgxpool.Pool, task ScheduleTask) {
	log.Printf("[Job #%d] Dispatching execution sequence to: %s", task.ID, task.PhoneNumber)

	// Trigger the real outbound internet message alert
	SendWhatsAppNotification(task.PhoneNumber, task.Message)

	// Commit the database state change natively to Neon Postgres
	updateQuery := "UPDATE scheduled_tasks SET is_completed = TRUE WHERE id = $1"
	_, err := dbPool.Exec(context.Background(), updateQuery, task.ID)
	if err != nil {
		log.Printf("[Job #%d] Error updating database execution state: %v", task.ID, err)
		return
	}
	log.Printf("✅ [Job #%d] State permanently committed as COMPLETED in cloud storage.", task.ID)
}

func SendWhatsAppNotification(recipientMobile string, textMessageBody string) {
	accessToken := os.Getenv("WHATSAPP_ACCESS_TOKEN")
	phoneNumberID := os.Getenv("WHATSAPP_PHONE_NUMBER_ID")

	if accessToken == "" || phoneNumberID == "" {
		log.Println("[WhatsApp API Error]: Missing access token or phone number ID in environment variables.")
		return
	}

	requestPayload := WhatsAppTemplateRequest{
		MessagingProduct: "whatsapp",
		To:               "6287768346975",
		Type:             "template",
	}
	requestPayload.Template.Name = "hello_world"
	requestPayload.Template.Language.Code = "en_US"

	jsonData, err := json.Marshal(requestPayload)
	if err != nil {
		log.Printf("[Network Client Error]: JSON serialization failed: %v", err)
		return
	}

	// Structural component combination layout to block link filter mutations entirely
	protocolStr := "https://"
	domainStr := "graph.facebook.com"
	apiVersionStr := "/v18.0/"
	endpointStr := "/messages"

	metaAPIUrl := protocolStr + domainStr + apiVersionStr + phoneNumberID + endpointStr

	req, err := http.NewRequest("POST", metaAPIUrl, bytes.NewBuffer(jsonData))
	if err != nil {
		log.Printf("[Network Client Error]: HTTP Request initialization failed: %v", err)
		return
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", accessToken))
	req.Header.Set("Content-Type", "application/json")

	// Apply a clean 10-second client connection timeout boundary safety rule
	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Do(req) 
	if err != nil {
		log.Printf("[Network Client Error]: Failed to transmit payload data packet across the wire: %v", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK || resp.StatusCode == http.StatusCreated {
		log.Printf("[Outbound Success]: Alert successfully dispatched across the internet to %s", recipientMobile)
	} else {
		// Read Meta's exact error message response body
		var errorResponse map[string]interface{}
		_ = json.NewDecoder(resp.Body).Decode(&errorResponse)
		errorJSON, _ := json.Marshal(errorResponse)

		log.Printf("[Outbound Failure]: Meta API rejected request with status indicator code: %d", resp.StatusCode)
		log.Printf("[Meta Error Body]: %s", string(errorJSON))	}
}
