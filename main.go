package main

import (
	"bank/api"
	"bank/controller"
	"bank/db"
	"bank/token"
	"bank/util"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

type server struct {
	tokenMaker token.Maker
}

func init() {
	db.ConnectDb()
}

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("Cannot load Config:", err)
	}
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("currency", controller.ValidCurrency)
	}

	server := gin.New()
	server.POST("/users/login", controller.LoginUser)
	server.POST("/createuser", controller.CreateUsers)
	authRoutes := server.Group("/").Use(api.AuthMiddleware())
	authRoutes.POST("/createaccount", controller.CreateAccount)
	authRoutes.GET("/accountdetails", controller.GetAccountDetails)
	authRoutes.GET("/accountdetails/:id", controller.FindAccountDetails)
	authRoutes.PATCH("/accountdetails/:id", controller.UpdateAccountDetails)
	authRoutes.DELETE("/accountdetails/:id", controller.AccountDetailsDelete)

	authRoutes.POST("/createtransfer", controller.CreateTransfer)
	authRoutes.POST("/createentry", controller.CreateEntry)
	authRoutes.GET("/entrydetails", controller.GetEntries)

	server.Run(config.Serverport)

}
