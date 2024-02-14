package controllers

import (
	"fmt"
	"net/http"
	"trab02/database"
	"trab02/models"
	"trab02/rabbitMQ"
	"trab02/service01/repositories"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	var (
		body      models.AuthDto
		userQuery models.User
	)

	db, err := database.Connect()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Resposta": err.Error()})
		return
	}

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, fmt.Sprintf("Resposta: %v", err.Error()))
		return
	}

	repo := repositories.NewUserRepository(db)
	userQuery, err = repo.SearchByEmail(body.Email)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"Resposta": err.Error()})
		return
	}

	if body.Password != userQuery.Password {
		c.JSON(http.StatusUnauthorized, gin.H{"Resposta": "Falha de autentificação de senha"})
		return
	}

	token, err := rabbitMQ.SendTokenGenerationRequest(userQuery.Id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Resposta": err.Error()})
		return
	}
	c.String(http.StatusOK, string(token))
}
