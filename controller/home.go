package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetHome(g *gin.Context) {
	g.IndentedJSON(http.StatusOK, gin.H{
		"message": "Hello there! <3",
	})
}
