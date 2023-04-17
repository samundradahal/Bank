package controller

import (
	"bank/db"
	"bank/models"
	"bank/token"
	"bank/util"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type userResponse struct {
	Username          string    `json:"username" `
	Email             string    `json:"email" `
	FullName          string    `json:"full_name" `
	PasswordChangedAt time.Time `json:password_changed_at`
	CreatedAt         time.Time `json:created_at`
}

func newUserResponse(user models.Users) userResponse {
	return userResponse{
		Username:          user.Username,
		Email:             user.Email,
		FullName:          user.FullName,
		PasswordChangedAt: user.PasswordChangedAt,
		CreatedAt:         user.CreatedAt,
	}
}

func CreateUsers(ctx *gin.Context) {

	var body struct {
		Username string `json:"username" binding:"required,alphanum"`
		Password string `json:"password" binding:"required,min=6"`
		Email    string `json:"email" binding:"required,email"`
		FullName string `json:"full_name" binding:"required"`
	}

	err := ctx.ShouldBind(&body)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}
	hashedPassword, err := util.HashPassword(body.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	user := models.Users{Username: body.Username, HashedPassword: hashedPassword, Email: body.Email, FullName: body.FullName}

	result := db.DB.Create(&user)
	if result.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": result.Error})
		return
	}
	var rsp = newUserResponse(user)

	ctx.JSON(http.StatusOK, gin.H{"message": "Details Added Successfully", "User Details": rsp})

}

type loginUserRequest struct {
	Username string `json:"username" binding:"required,alphanum"`
	Password string `json:"password" binding:"required,min=6"`
}

type loginUserResponse struct {
	AccessToken string       `json:"access_token"`
	User        userResponse `json:"user"`
}

func LoginUser(ctx *gin.Context) {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("Cannot load Config:", err)
	}
	var detail models.Users
	var req loginUserRequest
	err = ctx.ShouldBind(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	if err := db.DB.Where("username = ?", req.Username).First(&detail).Error; err != nil {
		ctx.JSON(http.StatusNotFound, err.Error())
		return
	}

	err = util.CheckPassword(req.Password, detail.HashedPassword)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, err.Error())
		return
	}

	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	accessToken, err := tokenMaker.CreateToken(
		detail.Username,
		config.AccessTokenDuration,
	)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	rsp := loginUserResponse{
		AccessToken: accessToken,
		User:        newUserResponse(detail),
	}
	ctx.JSON(http.StatusOK, gin.H{"Message": rsp})

}
