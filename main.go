package main

import (
	"fredmnpinto/DmLittleHelper/controller"
	"fredmnpinto/DmLittleHelper/service"
	"github.com/gin-gonic/gin"
)

func main() {
	service.SetupDices()

	r := gin.Default()

	r.GET("/", controller.GetHome)
	r.GET("/roll/:sides", controller.GetDiceRoll)

	err := r.Run("0.0.0.0:80")
	if err != nil {
		panic(err)
	}
}
