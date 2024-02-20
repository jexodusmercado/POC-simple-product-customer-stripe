package api

import (
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jexodusmercado/POC-simple-product-customer-stripe/internal/models"
)

func (api *API) CreateApplicant(ctx *gin.Context) {
	err := ctx.Request.ParseMultipartForm(10 << 20)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	applicant := models.ApplicantRequest{
		FirstName:       ctx.Request.FormValue("first_name"),
		LastName:        ctx.Request.FormValue("last_name"),
		Email:           ctx.Request.FormValue("email"),
		Phone:           ctx.Request.FormValue("phone"),
		ZipCode:         ctx.Request.FormValue("zip_code"),
		JobTitle:        ctx.Request.FormValue("job_title"),
		LinkedInProfile: ctx.Request.FormValue("linked_in_profile"),
	}
	// Parse ApplicationDate
	applicationDateStr := ctx.Request.FormValue("application_date")
	applicationDate, err := time.Parse("2006-01-02", applicationDateStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid application date format"})
		return
	}
	applicant.ApplicationDate = &applicationDate

	// Handle file attachment
	file, header, err := ctx.Request.FormFile("applicant_attachment")
	if err == nil {
		defer file.Close()

		// Read the file content
		fileContent, err := io.ReadAll(file)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read file content"})
			return
		}

		// Set the file content and filename in the applicant object
		applicant.ApplicantAttachment = fileContent
		applicant.ApplicantFileName = header.Filename
	} else if err != http.ErrMissingFile {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get file attachment"})
		return
	}

	createErr := models.CreateApplicant(api.db, applicant)
	if createErr != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	applicantReq := Applicant{
		FirstName:       applicant.FirstName,
		LastName:        applicant.LastName,
		Email:           applicant.Email,
		Phone:           applicant.Phone,
		ZipCode:         applicant.ZipCode,
		LinkedInProfile: applicant.LinkedInProfile,
		JobTitle:        applicant.JobTitle,
		ApplicationDate: applicant.ApplicationDate,
		ApplicantAttachment: applicant.ApplicantAttachment,
		ApplicantFileName: applicant.ApplicantFileName,
	}

	emailErr := api.SendApplicationMail(applicantReq)

	if emailErr != nil {
		errorMessage := fmt.Sprintf("Error sending contact us email: %v", err)
		ctx.JSON(http.StatusNotFound, gin.H{"error": errorMessage})
		return
	}

	// insert send email here to James@elated.io for prod / use your own email for dev

	ctx.JSON(http.StatusOK, gin.H{"message": "Applicant created successfully"})
}

func (api *API) GetAllApplicants(ctx *gin.Context) {
	applicants, err := models.GetAllApplicants(api.db)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, applicants)
}

func (api *API) GetApplicantById(ctx *gin.Context) {
	id := ctx.Param("id")
	applicant, err := models.GetApplicantById(api.db, id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, applicant)
}
