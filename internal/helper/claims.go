package helper

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type AuthClaims struct {
	UserID     uuid.UUID
	BusinessID uuid.UUID
	RoleID     uuid.UUID
	FirstName  string
	LastName   string
	Email      string
	Phone      string
}

func GetAuthClaims(ctx *gin.Context) *AuthClaims {
	userId, ok := ctx.Get("UserID")
	if !ok {
		return nil
	}

	businessId, ok := ctx.Get("BusinessID")
	if !ok {
		return nil
	}

	roleId, ok := ctx.Get("RoleID")
	if !ok {
		return nil
	}

	email, ok := ctx.Get("Email")
	if !ok {
		return nil
	}

	firstName, ok := ctx.Get("FirstName")
	if !ok {
		return nil
	}

	lastName, ok := ctx.Get("LastName")
	if !ok {
		return nil
	}

	phone, ok := ctx.Get("Phone")
	if !ok {
		return nil
	}

	userId, err := uuid.Parse(userId.(string))
	if err != nil {
		return nil
	}

	businessId, err = uuid.Parse(businessId.(string))
	if err != nil {
		return nil
	}

	roleId, err = uuid.Parse(roleId.(string))
	if err != nil {
		return nil
	}

	fmt.Println("CLAIMS")
	fmt.Println("userId", userId.(uuid.UUID))
	fmt.Println("companyId", businessId.(uuid.UUID))
	fmt.Println("roleId", roleId.(uuid.UUID))
	fmt.Println("email", email)
	fmt.Println("firstName", firstName)
	fmt.Println("lastName", lastName)

	return &AuthClaims{
		UserID:     userId.(uuid.UUID),
		BusinessID: businessId.(uuid.UUID),
		RoleID:     roleId.(uuid.UUID),
		FirstName:  firstName.(string),
		LastName:   lastName.(string),
		Phone:      phone.(string),
		Email:      email.(string),
	}

}
