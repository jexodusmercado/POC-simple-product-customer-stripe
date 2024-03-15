package api

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jexodusmercado/POC-simple-product-customer-stripe/internal/helper"
	"github.com/jexodusmercado/POC-simple-product-customer-stripe/internal/models"
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
	FirstName           string
	LastName            string
	Email               string
	Phone               string
	ZipCode             string
	LinkedInProfile     string
	JobTitle            string
	ApplicationDate     *time.Time
	ApplicantAttachment []byte
	ApplicantFileName   string
	FileUrl    string
}

type BetaList struct {
	FirstName   string
	LastName    string
	Email       string
	PhoneNumber string
	ZipCode     string
	IsJoinBeta  *time.Time
}

type ContactUs struct {
	FirstName   string
	LastName    string
	Email       string
	PhoneNumber string
	Message     string
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

	content := mail.NewContent("text/html", bodyContent.String())
	message := mail.NewV3MailInit(from, subject, to, content)

	sg := sendgrid.NewSendClient(api.config.SENDGRID_API_KEY)

	if len(req.ApplicantAttachment) > 0 {
		attachment := mail.NewAttachment()
		attachment.SetContent(base64.StdEncoding.EncodeToString(req.ApplicantAttachment))
		attachment.SetFilename(req.ApplicantFileName)

		if strings.HasSuffix(req.ApplicantFileName, ".pdf") {
			attachment.SetType("application/pdf")
		} else if strings.HasSuffix(req.ApplicantFileName, ".docx") {
			attachment.SetType("application/vnd.openxmlformats-officedocument.wordprocessingml.document")
		} else {
			fmt.Println("Unsupported file type:", req.ApplicantFileName)
		}

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

func (api *API) SendApplicationElatedMail(req Applicant) error {
	from := mail.NewEmail("info@elated.io", api.config.SENDGRID_EMAIL_FROM)
	to := mail.NewEmail(req.FirstName+" "+req.LastName, "info@elated.io")
	subject := "Application Submission - " + req.FirstName+" "+req.LastName + " for " + req.JobTitle + " Position"
	fmt.Println("Sending Application Elated email")

	executablePath, err := os.Getwd()
	if err != nil {
		fmt.Println("err. (1)")

		return err
	}

	templatesPath := filepath.Join(executablePath, "internal/templates")

	// Read HTML content template from file
	templatePath := filepath.Join(templatesPath, "applicant-elated-copy.html")
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

	content := mail.NewContent("text/html", bodyContent.String())
	message := mail.NewV3MailInit(from, subject, to, content)

	sg := sendgrid.NewSendClient(api.config.SENDGRID_API_KEY)

	if len(req.ApplicantAttachment) > 0 {
		attachment := mail.NewAttachment()
		attachment.SetContent(base64.StdEncoding.EncodeToString(req.ApplicantAttachment))
		attachment.SetFilename(req.ApplicantFileName)

		if strings.HasSuffix(req.ApplicantFileName, ".pdf") {
			attachment.SetType("application/pdf")
		} else if strings.HasSuffix(req.ApplicantFileName, ".docx") {
			attachment.SetType("application/vnd.openxmlformats-officedocument.wordprocessingml.document")
		} else {
			fmt.Println("Unsupported file type:", req.ApplicantFileName)
		}

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
	to := mail.NewEmail(req.FirstName+" "+req.LastName, "info@elated.io")
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

func (api *API) SendQrCodeMail(db *gorm.DB, c *gin.Context, user models.User, transaction models.Transaction, product models.Product) error {

	path := "internal/templates"

	if api.config.TEMPLATE_PATH != "" {
		path = api.config.TEMPLATE_PATH
	}

	from := mail.NewEmail("info@elated.io", api.config.SENDGRID_EMAIL_FROM)
	to := mail.NewEmail(user.FirstName+" "+user.LastName, user.Email)
	subject := "SparksFlirt Purchase Confirmation"

	executablePath, err := os.Getwd()
	if err != nil {
		fmt.Println("err. (1)")

		return err
	}

	templatesPath := filepath.Join(executablePath, path)

	fmt.Println("templatesPath: ", templatesPath)

	currentTime := time.Now()
	currentDateString := currentTime.Format("2006-01-02 15:04:05")

	var isEarlyAccess bool

	if user.IsJoinBeta != nil && !user.IsJoinBeta.IsZero() {
    	isEarlyAccess = true
	} else {
    	isEarlyAccess = false
	}

	qrCodeDetails := helper.QRCodeDetails{
		UserID:            user.ID.String(),
		ProductID:         product.ID.String(),
		TransactionID:     transaction.ID.String(),
		TransactionType:   "PURCHASE.SPARKFLIRT",
		UserName:          user.FirstName + " " + user.LastName,
		UserEmail:         user.Email,
		UserZipCode:       user.ZipCode,
		IsUserEarlyAccess: isEarlyAccess,
		ProductName:       product.Name,
		Description:       product.Description,
		Package:           strconv.Itoa(product.Quantity),
		BasePrice:         strconv.FormatFloat(product.BasePrice, 'f', -1, 64),
		PriceWithDiscount: strconv.FormatFloat(product.DiscountedPrice, 'f', -1, 64),
		Date:              currentDateString,
	}

	// Load the small image from its URL or any other source
	smallImageURL := "https://i.ibb.co/kKMwRpR/Group-988759.png"
	resp, err := http.Get(smallImageURL)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error getting small image",
		})
	}
	defer resp.Body.Close()
	smallImage, err := io.ReadAll(resp.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error generating QR code",
		})
	}

	qrCode, generateQrerr := helper.GenerateQRCodeWithEncryptedData(qrCodeDetails, []byte(api.config.KEY_ELATED), []byte(api.config.IV_ELATED), smallImage)
	if generateQrerr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error generating QR code: " + generateQrerr.Error(),
		})
	}

	key, objectUrl, qrErr := api.UploadQRCode(qrCode, user.ID.String())

	if qrErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error uploading QR code",
			"error":   qrErr,
		})
	}

	qrStorageReq := models.CreateQrCodeRequest{
		TransactionID: transaction.ID,
		UserID:        user.ID,
		S3Url:         objectUrl,
	}

	if _, err := models.CreateQrCode(api.db, &qrStorageReq); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error creating QR code entry",
			"error":   err,
		})
		return err
	}

	c.JSON(http.StatusOK, gin.H{
		"message":   "QR code uploaded",
		"key":       key,
		"objectUrl": objectUrl,
	})

	templatePath := filepath.Join(templatesPath, "qrcode.html")
	tmpl, err := template.ParseFiles(templatePath)
	if err != nil {
		fmt.Println("err. (3)", err)

		return err
	}

	var bodyContent bytes.Buffer

	// Pass the objectUrl value into the template
	err = tmpl.Execute(&bodyContent, struct {
		ObjectUrl string
		FirstName string
	}{
		ObjectUrl: objectUrl,
		FirstName: user.FirstName,
	})
	fmt.Println("objectUrl: ", objectUrl)
	if err != nil {
		fmt.Println("Execute err", err)

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
