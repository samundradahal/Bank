package controller

import (
	"bank/db"
	"bank/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateTransfer(ctx *gin.Context) {

	var body struct {
		From_Accountid int64  `json:"from_id" binding:"required"`
		To_Accountid   int64  `json:"to_id" binding:"required"`
		Amount         int64  `json:"amount" binding:"required"`
		Currency       string `json:"currency" binding:"required,currency"`
	}
	var detail1 models.Account
	var detail2 models.Account
	err := ctx.Bind(&body)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"Message": "Enter Right Data"})
		return
	}
	if !validAccount(ctx, body.From_Accountid, body.Currency) {
		return
	}
	if !validAccount(ctx, body.To_Accountid, body.Currency) {
		return
	}
	if body.From_Accountid == body.To_Accountid {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Cannot Transfer to same account"})
		return
	}
	user := models.Transfers{From_Accountid: body.From_Accountid, To_Accountid: body.To_Accountid, Amount: body.Amount}
	result := db.DB.Create(&user)

	if result.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Account Number not matched"})
		return
	}
	db.DB.First(&detail1, body.From_Accountid)
	db.DB.First(&detail2, body.To_Accountid)
	db.TransferMoney(body.From_Accountid, body.To_Accountid, body.Amount)
	ctx.JSON(http.StatusOK, gin.H{
		"Transfer Details":     user,
		"From Account Details": detail1,
		"To Account Details":   detail2,
		"message":              "Transfer Done Successfully "})
	Entry(user.From_Accountid, user.Amount*-1)
	Entry(user.To_Accountid, user.Amount)

}

func validAccount(ctx *gin.Context, accountID int64, currency string) bool {
	var inp models.Account
	err := db.DB.First(&inp, accountID).Error
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return false
	}
	if inp.Currency != currency {
		err := fmt.Sprintf("account [%d] currency mismatched: %s vs %s", inp.Id, inp.Currency, currency)
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err})
		return false
	}
	return true
}
