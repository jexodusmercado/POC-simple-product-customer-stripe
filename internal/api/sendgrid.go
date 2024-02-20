package api

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"html/template"
	"os"
	"path/filepath"
	"time"

	"github.com/google/uuid"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
	"gorm.io/gorm"
)

type Mail struct {
	Subject       string
	CustomerEmail string
	CustomerName  string
	Body          string
}

type Applicant struct {
	JobTitle          string
	FirstName         string
	LastName          string
	Email             string
	Phone             string
	ZipCode           string
	LinkedInProfile   string
	ApplicationDate     *time.Time
	ApplicantAttachment []byte
	ApplicantFileName string
}

type BetaList struct {
	FirstName      string
	LastName       string
	Email          string
	PhoneNumber    string
	ZipCode        string
	IsJoinBeta     *time.Time
}

type ContactUs struct {
	FirstName      string
	LastName       string
	Email          string
	PhoneNumber    string
	Message        string
}

type QrCode struct {
	TransactionId      uuid.UUID
	ProductId          uuid.UUID
	UserId             uuid.UUID
	UserName           string
	ProductName        string
	Description        string
	Package        	   string
	PriceWithDiscount  string
	Date        	   string
	QrCodeImgTag       string
}

func (api *API) SendMail(req Mail) error {
	from := mail.NewEmail("info@elated.io", api.config.SENDGRID_EMAIL_FROM)
	to := mail.NewEmail(req.CustomerName, req.CustomerEmail)
	message := mail.NewSingleEmail(from, req.Subject, to, req.Body, req.Body)
	client := sendgrid.NewSendClient(api.config.SENDGRID_API_KEY)

	_, err := client.Send(message)
	if err != nil {
		return err
	}

	return nil
}

func (api *API) SendApplicationMail(req Applicant) error {
	from := mail.NewEmail("info@elated.io", api.config.SENDGRID_EMAIL_FROM)
	to := mail.NewEmail(req.FirstName+" "+req.LastName, req.Email)
	subject := "Application Acknowledgement - " + req.JobTitle + " Position"
	fmt.Println("Sending Application email")

	executablePath, err := os.Getwd()
	if err != nil {
		fmt.Println("err. (1)")

		return err
	}

	templatesPath := filepath.Join(executablePath, "internal/templates")

	// Read HTML content template from file
	templatePath := filepath.Join(templatesPath, "applicant.html")
	tmpl, err := template.ParseFiles(templatePath)
	if err != nil {
		fmt.Println("err. (3)", err)

		return err
	}

	var bodyContent bytes.Buffer

	err = tmpl.Execute(&bodyContent, req)
	if err != nil {
		fmt.Println("Execute err")

		return err
	}

	// Code to send email using SendGrid
	content := mail.NewContent("text/html", bodyContent.String())
	message := mail.NewV3MailInit(from, subject, to, content)

	// Initialize SendGrid client using the API key
	sg := sendgrid.NewSendClient(api.config.SENDGRID_API_KEY)

	if len(req.ApplicantAttachment) > 0 {
        attachment := mail.NewAttachment()
        attachment.SetContent(base64.StdEncoding.EncodeToString(req.ApplicantAttachment))
        attachment.SetType("application/pdf") // Set the attachment type (e.g., application/pdf)
        attachment.SetFilename(req.ApplicantFileName) // Set the filename

        // Add attachment to the message
        message.AddAttachment(attachment)
    }

	res, err := sg.Send(message)
	if err != nil {
		fmt.Println("err. (5)")

		return err
	}

	if res.StatusCode != 200 && res.StatusCode != 202 {
		fmt.Println("err. (6)")

		return err
	}
	fmt.Println("Sending Customer Email Success.")

	return nil
}

func (api *API) SendBetaMail(req BetaList) error {
	from := mail.NewEmail("info@elated.io", api.config.SENDGRID_EMAIL_FROM)
	to := mail.NewEmail(req.FirstName+" "+req.LastName, req.Email)
	subject := "Elated Beta List Registration Confirmation"
	fmt.Println("Sending Registration Confirmation email")

	executablePath, err := os.Getwd()
	if err != nil {
		fmt.Println("err. (1)")

		return err
	}

	templatesPath := filepath.Join(executablePath, "internal/templates")

	// Read HTML content template from file
	templatePath := filepath.Join(templatesPath, "beta.html")
	tmpl, err := template.ParseFiles(templatePath)
	if err != nil {
		fmt.Println("err. (3)", err)

		return err
	}

	var bodyContent bytes.Buffer

	err = tmpl.Execute(&bodyContent, req)
	if err != nil {
		fmt.Println("Execute err")

		return err
	}

	// Code to send email using SendGrid
	content := mail.NewContent("text/html", bodyContent.String())
	message := mail.NewV3MailInit(from, subject, to, content)

	// Initialize SendGrid client using the API key
	sg := sendgrid.NewSendClient(api.config.SENDGRID_API_KEY)

	res, err := sg.Send(message)
	if err != nil {
		fmt.Println("err. (5)")

		return err
	}

	if res.StatusCode != 200 && res.StatusCode != 202 {
		fmt.Println("err. (6)")

		return err
	}
	fmt.Println("Sending Beta List Confirmation Email Success.")

	return nil
}

func (api *API) SendContactUsMail(req ContactUs) error {
	from := mail.NewEmail("info@elated.io", api.config.SENDGRID_EMAIL_FROM)
	to := mail.NewEmail(req.FirstName+" "+req.LastName, req.Email)
	subject := "Elated Contact Us Inquiry"
	fmt.Println("Sending Inquiry email")

	executablePath, err := os.Getwd()
	if err != nil {
		fmt.Println("err. (1)")

		return err
	}

	templatesPath := filepath.Join(executablePath, "internal/templates")

	// Read HTML content template from file
	templatePath := filepath.Join(templatesPath, "inquiry.html")
	tmpl, err := template.ParseFiles(templatePath)
	if err != nil {
		fmt.Println("err. (3)", err)

		return err
	}

	var bodyContent bytes.Buffer

	err = tmpl.Execute(&bodyContent, req)
	if err != nil {
		fmt.Println("Execute err")

		return err
	}

	// Code to send email using SendGrid
	content := mail.NewContent("text/html", bodyContent.String())
	message := mail.NewV3MailInit(from, subject, to, content)

	// Initialize SendGrid client using the API key
	sg := sendgrid.NewSendClient(api.config.SENDGRID_API_KEY)

	res, err := sg.Send(message)
	if err != nil {
		fmt.Println("err. (5)")

		return err
	}

	if res.StatusCode != 200 && res.StatusCode != 202 {
		fmt.Println("err. (6)")

		return err
	}
	fmt.Println("Sending Beta List Confirmation Email Success.")

	return nil
}

func (api *API) SendQrCodeMail(db *gorm.DB, req QrCode) error {

	//user, err := models.GetUserByID(db, req.UserId)
	//if err != nil {
	//	return err
	//}

	from := mail.NewEmail("info@elated.io", api.config.SENDGRID_EMAIL_FROM)
	to := mail.NewEmail("blue", "blueandraedevera@gmail.com")
	subject := "SparksFlirt Purchase Confirmation"
	fmt.Println("Sending SparksFlirt Purchase Confirmation email")

	executablePath, err := os.Getwd()
	if err != nil {
		fmt.Println("err. (1)")

		return err
	}

	templatesPath := filepath.Join(executablePath, "internal/templates")
	qrCodeData := QrCode{
		QrCodeImgTag: `<img src="https://api.qrserver.com/v1/create-qr-code/?size=150x150&data=12332131231" alt="QR Code" style="margin-top: 20px;">`,
	}
	req.QrCodeImgTag = qrCodeData.QrCodeImgTag
	// Read HTML content template from file
	templatePath := filepath.Join(templatesPath, "qrcode.html")
	tmpl, err := template.ParseFiles(templatePath)
	if err != nil {
		fmt.Println("err. (3)", err)

		return err
	}

	var bodyContent bytes.Buffer
	
	err = tmpl.Execute(&bodyContent, req)
	if err != nil {
		fmt.Println("Execute err")

		return err
	}

	// Code to send email using SendGrid
	content := mail.NewContent("text/html", bodyContent.String())
	message := mail.NewV3MailInit(from, subject, to, content)
	
	// Initialize SendGrid client using the API key
	sg := sendgrid.NewSendClient(api.config.SENDGRID_API_KEY)

	res, err := sg.Send(message)
	if err != nil {
		fmt.Println("err. (5)")

		return err
	}

	if res.StatusCode != 200 && res.StatusCode != 202 {
		fmt.Println("err. (6)")

		return err
	}
	fmt.Println("Sending Customer Email Success.")

	return nil
}