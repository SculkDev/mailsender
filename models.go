package main

type TemplateType string

var ignoredTemplateTypes = map[TemplateType]bool{
	"recovery_invalid":          true,
	"recovery_code_invalid":     true,
	"verification_invalid":      true,
	"verification_code_invalid": true,
}

type MailRequest struct {
	Body             *string `json:"body"`
	LoginCode        *string `json:"login_code"`
	Recipient        *string `json:"recipient"`
	RecoveryCode     *string `json:"recovery_code"`
	RecoveryURL      *string `json:"recovery_url"`
	RegistrationCode *string `json:"registration_code"`
	Subject          *string `json:"subject"`
	Template         *string `json:"template_type"`
	To               *string `json:"to"`
	VerificationCode *string `json:"verification_code"`
	VerificationURL  *string `json:"verification_url"`
}

func (r *MailRequest) TemplateVars() map[string]string {
	vars := make(map[string]string)
	if r.Body != nil {
		vars["body"] = *r.Body
	}
	if r.LoginCode != nil {
		vars["login_code"] = *r.LoginCode
	}
	if r.Recipient != nil {
		vars["recipient"] = *r.Recipient
	}
	if r.RecoveryCode != nil {
		vars["recovery_code"] = *r.RecoveryCode
	}
	if r.RecoveryURL != nil {
		vars["recovery_url"] = *r.RecoveryURL
	}
	if r.RegistrationCode != nil {
		vars["registration_code"] = *r.RegistrationCode
	}
	if r.Subject != nil {
		vars["subject"] = *r.Subject
	}
	if r.To != nil {
		vars["to"] = *r.To
	}
	if r.VerificationCode != nil {
		vars["verification_code"] = *r.VerificationCode
	}
	if r.VerificationURL != nil {
		vars["verification_url"] = *r.VerificationURL
	}
	return vars
}
