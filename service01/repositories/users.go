package repositories

import (
	"trab02/models"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) GetUsers() ([]models.UserResponse, error) {
	var users []models.User
	result := r.db.Find(&users)
	userResponses := models.ToUserResponse(users)
	return userResponses, result.Error
}

func (r *UserRepository) GetUser(id uint64) (models.UserResponse, error) {
	var user models.User
	result := r.db.First(&user, id)
	userResponse := models.ToUserResponse([]models.User{user})[0]
	return userResponse, result.Error
}

func (r *UserRepository) PostUser(newUser models.User) (uint64, error) {
	result := r.db.Create(&newUser)
	if result.Error != nil {
		return 0, result.Error
	}
	return newUser.Id, nil
}
