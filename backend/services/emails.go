package services

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-chi/render"
	"github.com/walteralv/enronMailApp/adapter"
)

// Email is the struct for the email
type Email struct {
	MessageID string `json:"message_id"`
	Date      string `json:"date"`
	From      string `json:"from"`
	To        string `json:"to"`
	Subject   string `json:"subject"`
	Content   string `json:"content"`
}

type SearchEmailsResponse struct {
	Emails []Email `json:"emails"`
}

func mapZincSearchResponseToEmails(response *adapter.SearchDocumentsResponse) []Email {
	var emails []Email

	for _, hit := range response.Hits.Hits {
		var email Email

		emailBytes, _ := json.Marshal(hit.Source)

		err := json.Unmarshal(emailBytes, &email)
		if err != nil {
			log.Println("Error in mapZincSearchResponseToEmails: ", err)
			continue
		}

		emails = append(emails, email)
	}

	return emails
}

func SearchEmail(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	term := query.Get("q")

	records, err := adapter.SearchDocuments("emails", term)
	if err != nil {
		panic(err)
	}

	response := &SearchEmailsResponse{
		Emails: mapZincSearchResponseToEmails(records),
	}

	render.JSON(w, r, response)
}
