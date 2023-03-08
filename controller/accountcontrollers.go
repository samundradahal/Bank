package controller

import (
	"bank/db"
	"bank/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateAccount(ctx *gin.Context) {

	var body struct {
		Owner    string `json:"owner" binding:"required"`
		Currency string `json:"currency" binding:"required,currency"`
	}

	err := ctx.ShouldBind(&body)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"Message": "Enter Right Data"})
		return
	}
	user := models.Account{Owner: body.Owner, Balance: 0, Currency: body.Currency}

	result := db.DB.Create(&user)
	if result.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Error creating post"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Details Added Successfully", "Account Details": user})

}

func GetAccountDetails(ctx *gin.Context) {
	var detail []models.Account
	db.DB.Find(&detail)
	ctx.JSON(http.StatusOK, gin.H{
		"message": "All Account Details",
		"Details": detail,
	})

}

func FindAccountDetails(ctx *gin.Context) {
	var detail models.Account
	id := ctx.Param("id")

	db.DB.First(&detail, id)
	err := db.DB.First(&detail, id).Error

	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"Message": "Account details not found",
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"Message": "Account details found",
		"Post":    detail,
	})
}

func UpdateAccountDetails(ctx *gin.Context) {
	id := ctx.Param("id")
	var input struct {
		Owner    string `json:"owner" binding:"required"`
		Currency string `json:"currency" binding:"required,oneof=NPR USD INR"`
	}

	ctx.Bind(&input)

	var det models.Account

	if err := db.DB.First(&det, id).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"Error": "Account Details not found"})
		return
	}

	if err := db.DB.Model(&det).Updates(&models.Account{Owner: input.Owner, Currency: input.Currency}).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update details"})
		return
	}

	s := fmt.Sprintf("Details updated of Account id = %v", id)
	ctx.JSON(http.StatusOK, gin.H{"message": s, "New Details": det})
}

func AccountDetailsDelete(ctx *gin.Context) {
	id := ctx.Param("id")

	if (db.DB.First(&models.Account{}, id).RowsAffected == 0) {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Account Details not found"})
		return
	}

	db.DB.Delete(&models.Account{}, id)
	s := fmt.Sprintf("Post deleted successfully of Account  id %v", id)
	ctx.JSON(http.StatusOK, gin.H{"Message": s})
}
