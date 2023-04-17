package main

import (
	"bank/controller"
	"bank/db"
	"bank/util"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

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
	server.POST("/createaccount", controller.CreateAccount)
	server.GET("/accountdetails", controller.GetAccountDetails)
	server.GET("/accountdetails/:id", controller.FindAccountDetails)
	server.PATCH("/accountdetails/:id", controller.UpdateAccountDetails)
	server.DELETE("/accountdetails/:id", controller.AccountDetailsDelete)

	server.POST("/createtransfer", controller.CreateTransfer)
	server.POST("/createentry", controller.CreateEntry)
	server.GET("/entrydetails", controller.GetEntries)

	server.POST("/users/login", controller.LoginUser)
	server.POST("/createuser", controller.CreateUsers)
	server.Run(config.Serverport)

}
