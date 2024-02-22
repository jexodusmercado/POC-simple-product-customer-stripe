package api

import (
	"fmt"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jexodusmercado/POC-simple-product-customer-stripe/internal/models"
)

func (api *API) CreateApplicant(ctx *gin.Context) {
	var applicant models.ApplicantRequest

	if err := ctx.ShouldBind(&applicant); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	file, err := ctx.FormFile("applicant_attachment")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	fmt.Println(file.Filename)

	applicant.FileURL, err = api.UploadAttachment(file)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	err = models.CreateApplicant(api.db, applicant)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	fileContent, err := file.Open()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer fileContent.Close()

	fileBytes, err := io.ReadAll(fileContent)

	applicantReq := Applicant{
		FirstName:         applicant.FirstName,
		LastName:          applicant.LastName,
		Email:             applicant.Email,
		Phone:             applicant.Phone,
		ZipCode:           applicant.ZipCode,
		LinkedInProfile:   applicant.LinkedInProfile,
		JobTitle:          applicant.JobTitle,
		ApplicationDate:   applicant.ApplicationDate,
		ApplicantFileName: file.Filename,
		ApplicantAttachment: fileBytes,
	}

	emailErr := api.SendApplicationMail(applicantReq)

	if emailErr != nil {
		errorMessage := fmt.Sprintf("Error sending application us email: %v", err)
		ctx.JSON(http.StatusNotFound, gin.H{"error": errorMessage})
		return
	}

	emailElatedErr := api.SendApplicationElatedMail(applicantReq)

	if emailElatedErr != nil {
		errorMessage := fmt.Sprintf("Error sending elated application us email: %v", err)
		ctx.JSON(http.StatusNotFound, gin.H{"error": errorMessage})
		return
	}

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
