package controllers

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"trab02/database"
	"trab02/models"
	"trab02/service01/repositories"

	"github.com/gin-gonic/gin"
)

func GetUsers(c *gin.Context) {
	db, err := database.Connect()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Resposta": err.Error()})
		return
	}

	repo := repositories.NewUserRepository(db)
	users, err := repo.GetUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Resposta": err.Error()})
		return
	}

	c.JSON(http.StatusOK, users)
}

func GetUser(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Resposta": err.Error()})
		return
	}

	db, err := database.Connect()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Resposta": err.Error()})
		return
	}
	repo := repositories.NewUserRepository(db)
	user, err := repo.GetUser(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"Resposta": err.Error()})
		return
	}
	c.JSON(http.StatusOK, user)
}

func PostUser(c *gin.Context) {
	var newUser models.User

	db, err := database.Connect()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Resposta": err.Error()})
		return
	}

	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, fmt.Sprintf("Erro: %v", err.Error()))
		return
	}

	if newUser.Name == "" || newUser.Email == "" || newUser.Password == "" {
		c.JSON(http.StatusBadRequest, gin.H{"Resposta": "Algum dado do body está vazio"})
		return
	}

	repo := repositories.NewUserRepository(db)
	newInsertedId, err := repo.PostUser(newUser)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Resposta": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"Resposta": fmt.Sprintf("Novo usuário inserido com id %v", newInsertedId)})
}

func DeleteUser(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Resposta": err.Error()})
		return
	}

	// validating the token
	if err := validateToken(c, id); err != nil {
		c.JSON(http.StatusUnauthorized, fmt.Sprintf("Erro: %v", err.Error()))
		return
	}

	db, err := database.Connect()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Resposta": err.Error()})
		return
	}

	repo := repositories.NewUserRepository(db)
	err = repo.DeleteUser(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Resposta": err.Error()})
		return
	}
	c.JSON(http.StatusOK, fmt.Sprintf("Usuario de id %v deletado", id))
}

func UpdateUser(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Resposta": err.Error()})
		return
	}

	if err := validateToken(c, id); err != nil {
		c.JSON(http.StatusUnauthorized, fmt.Sprintf("Erro: %v", err.Error()))
		return
	}

	var newUserData models.User
	newUserData.Id = id
	if err := c.ShouldBindJSON(&newUserData); err != nil {
		c.JSON(http.StatusBadRequest, fmt.Sprintf("Erro: %v", err.Error()))
		return
	}

	db, err := database.Connect()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Resposta": err.Error()})
		return
	}

	repo := repositories.NewUserRepository(db)
	err = repo.UpdateUser(newUserData)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Resposta": err.Error()})
		return
	}
	c.JSON(http.StatusOK, fmt.Sprintf("Usuário de id %v atualizado", id))
}

func validateToken(c *gin.Context, id uint64) error {
	header := strings.Split(c.GetHeader("Authorization"), " ")
	if len(header) < 2 {
		return errors.New("token em branco")
	}
	bodyToken := header[1]

	request, err := http.NewRequest("POST", fmt.Sprintf("http://localhost:8081/usuarios/validar-token?userID=%d", id), nil)
	if err != nil {
		return errors.New("erro ao criar a requisição")
	}

	request.Header.Set("Authorization", "Bearer "+bodyToken)
	client := &http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		return errors.New("erro ao executar a requisição")
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusUnauthorized {
		return errors.New("requisição não autorizada")
	}
	return nil
}
