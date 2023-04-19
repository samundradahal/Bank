package controller

import (
	"bank/db"
	"bank/models"
	"bank/token"
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	authorizationPayloadKey = "authorization_payload"
)

func CreateAccount(ctx *gin.Context) {

	var body struct {
		Currency string `json:"currency" binding:"required,currency"`
	}

	err := ctx.ShouldBind(&body)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"Message": "Enter Right Data"})
		return
	}
	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)
	user := models.Account{Owner: authPayload.Username, Balance: 0, Currency: body.Currency}

	result := db.DB.Create(&user)
	if result.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Error creating post"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Details Added Successfully", "Account Details": user})

}

func GetAccountDetails(ctx *gin.Context) {
	var detail []models.Account
	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)

	db.DB.Where("owner <> ?", authPayload.Username).Find(&detail)
	ctx.JSON(http.StatusOK, gin.H{
		"message": "All Account Details",
		"Details": detail,
	})

}

func FindAccountDetails(ctx *gin.Context) {
	var detail models.Account
	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)
	id := ctx.Param("id")

	db.DB.First(&detail, id)
	err := db.DB.First(&detail, id).Error

	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"Message": "Account details not found",
		})
		return
	}
	if detail.Owner != authPayload.Username {
		err := errors.New("account doesn't belong to authorized user")
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"Message": "Account details found",
		"Post":    detail,
	})
}

func UpdateAccountDetails(ctx *gin.Context) {
	id := ctx.Param("id")
	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)
	var input struct {
		Currency string `json:"currency" binding:"required,oneof=NPR USD INR"`
	}

	ctx.Bind(&input)

	var det models.Account
	if det.Owner != authPayload.Username {
		err := errors.New("account doesn't belong to authorized user")
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": err.Error()})
		return
	}

	if err := db.DB.First(&det, id).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"Error": "Account Details not found"})
		return
	}

	if err := db.DB.Model(&det).Updates(&models.Account{Currency: input.Currency}).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update details"})
		return
	}

	s := fmt.Sprintf("Details updated of Account id = %v", id)
	ctx.JSON(http.StatusOK, gin.H{"message": s, "New Details": det})
}

func AccountDetailsDelete(ctx *gin.Context) {
	id := ctx.Param("id")
	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)
	fmt.Println(authPayload)
	if (db.DB.First(&models.Account{}, id).RowsAffected == 0) {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Account Details not found"})
		return
	}

	db.DB.Delete(&models.Account{}, id)
	s := fmt.Sprintf("Post deleted successfully of Account  id %v", id)
	ctx.JSON(http.StatusOK, gin.H{"Message": s})
}
