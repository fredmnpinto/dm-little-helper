package controller

import (
	"fredmnpinto/DmLittleHelper/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func GetDiceRoll(g *gin.Context) {
	sParam := g.Param("sides")
	n, parseErr := strconv.ParseInt(sParam, 10, 32)

	if parseErr != nil {
		g.IndentedJSON(http.StatusBadRequest, gin.H{
			"message": "you sure you got the right number? '" + sParam + "' looks weird.",
		})

	}

	roll := service.RollDice(int(n))

	g.IndentedJSON(http.StatusOK, gin.H{
		"result": strconv.FormatInt(int64(roll), 10),
	})

}
