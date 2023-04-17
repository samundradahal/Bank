package controller

import (
	"bank/db"
	"bank/models"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Entry(accountid int64, amount int64) {
	user := models.Entries{AccountId: accountid, Amount: amount}
	result := db.DB.Create(&user)
	if result.Error != nil {
		log.Fatal("Error to create error")
		return
	}
}

func CreateEntry(ctx *gin.Context) {

	var body struct {
		AccountId int64 `json:"accountid" binding:"required"`
		Amount    int64 `json:"amount" binding:"required"`
	}

	err := ctx.ShouldBind(&body)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"Message": "Enter Right Data"})
		return
	}

	user := models.Entries{AccountId: body.AccountId, Amount: body.Amount}
	result := db.DB.Create(&user)
	if result.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Error creating post"})
		return
	}
	db.UpdateAccount(user.AccountId, user.Amount)
	ctx.JSON(http.StatusOK, gin.H{"message": "Entry Added Successfully", "Entry": user})

}

func GetEntries(ctx *gin.Context) {
	var detail []models.Entries
	db.DB.Find(&detail)
	ctx.JSON(http.StatusOK, gin.H{
		"message": "All Entries Details",
		"Details": detail,
	})

}
