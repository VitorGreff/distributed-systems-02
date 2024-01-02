package main

import (
	"net/http"
	"strconv"
	"trab02/service02/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/usuarios/token", func(c *gin.Context) {
		userID, err := strconv.ParseUint(c.Query("userID"), 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"Resposta": "id fornecido inválido"})
			return
		}
		token, err := controllers.GenerateToken(c, userID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"Resposta": err.Error()})
			return
		}
		c.String(http.StatusOK, token)
	})
	router.POST("/usuarios/validar-token", func(c *gin.Context) {
		userID, err := strconv.ParseUint(c.Query("userID"), 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"Resposta": "id fornecido inválido"})
			return
		}

		if err := controllers.ValidateToken(c, userID); err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"Resposta": err.Error()})
			return
		}
		c.JSON(http.StatusAccepted, gin.H{"Resposta": "Token válido"})
	})
	router.Run(":8081")
}
