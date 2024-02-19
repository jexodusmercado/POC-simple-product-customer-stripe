package api

import (
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

type Mail struct {
	Subject       string
	CustomerEmail string
	CustomerName  string
	Body          string
}

func (api *API) SendMail(req Mail) error {
	from := mail.NewEmail("Elated", api.config.SENDGRID_EMAIL_FROM)
	to := mail.NewEmail(req.CustomerName, req.CustomerEmail)
	message := mail.NewSingleEmail(from, req.Subject, to, req.Body, req.Body)
	client := sendgrid.NewSendClient(api.config.SENDGRID_API_KEY)

	_, err := client.Send(message)
	if err != nil {
		return err
	}

	return nil
}
