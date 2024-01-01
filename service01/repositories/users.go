package repositories

import (
	"errors"
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

func (r *UserRepository) DeleteUser(id uint64) error {
	user, err := r.GetUser(id)
	if err != nil || user.Id == 0 {
		return errors.New("usuário não está cadastrado no banco")
	}

	result := r.db.Delete(&models.User{}, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *UserRepository) SearchByEmail(email string) (models.User, error) {
	var user models.User

	r.db.Raw("select id, password from users where email = ?", email).Scan(&user)
	if user.Id == 0 {
		return models.User{}, errors.New("usuário não está cadastrado no banco")
	}

	return user, nil
}
