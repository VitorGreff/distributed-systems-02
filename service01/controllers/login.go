package controllers

import (
	"fmt"
	"io"
	"net/http"
	"trab02/database"
	"trab02/models"
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
		c.JSON(http.StatusBadRequest, gin.H{"Resposta": err.Error()})
		return
	}

	if body.Password != userQuery.Password {
		c.JSON(http.StatusUnauthorized, gin.H{"Resposta": "Falha de autentificação de senha"})
		return
	}

	// request for getting the token
	response, err := http.Get(fmt.Sprintf("http://localhost:8081/usuarios/token?userID=%d", userQuery.Id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Resposta": err.Error()})
		return
	}

	token, err := io.ReadAll(response.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Resposta": err.Error()})
		return
	}

	c.String(http.StatusOK, string(token))
}
