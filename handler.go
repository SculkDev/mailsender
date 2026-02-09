package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/getsentry/sentry-go"
)

func mailhookHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	internalKey := os.Getenv("MAILSENDER_INTERNAL_KEY")
	if internalKey != "" && r.Header.Get("X-Internal-Key") != internalKey {
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return
	}

	var req MailRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	if req.Template == nil {
		http.Error(w, "template_type is required", http.StatusBadRequest)
		return
	}

	tt := TemplateType(*req.Template)

	if ignoredTemplateTypes[tt] {
		w.WriteHeader(http.StatusOK)
		return
	}

	to := ""
	if req.To != nil {
		to = *req.To
	} else if req.Recipient != nil {
		to = *req.Recipient
	}
	if to == "" {
		http.Error(w, "recipient (to) is required", http.StatusBadRequest)
		return
	}

	vars := req.TemplateVars()

	htmlBody, textBody, subject, err := RenderTemplate(tt, vars)
	if err != nil {
		log.Printf("template render error: %v", err)
		sentry.CaptureException(err)
		http.Error(w, "template render error", http.StatusInternalServerError)
		return
	}

	if req.Subject != nil {
		subject = *req.Subject
	}

	cfg := SMTPConfigFromEnv()
	if err := SendMail(to, cfg.From, subject, htmlBody, textBody, cfg); err != nil {
		log.Printf("smtp send error: %v", err)
		sentry.CaptureException(err)
		http.Error(w, "failed to send email", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
