package auth

import (
	"errors"

	"example.com/m/v2/models"
	"gorm.io/gorm"
)

type AuthRepository struct {
	db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) RepositoryAuthInterface {
	return &AuthRepository{db: db}
}

// RegisterUser - Menyimpan user baru ke database
func (r *AuthRepository) RegisterUser(newUser *models.UserModel) (*models.UserModel, error) {
	if err := r.db.Create(newUser).Error; err != nil {
		return nil, errors.New("failed to register user")
	}
	return newUser, nil
}

// LoginUser - Mencari user berdasarkan username
func (r *AuthRepository) LoginUser(username, password string) (*models.UserModel, error) {
	var user models.UserModel
	if err := r.db.Where("username = ?", username).First(&user).Error; err != nil {
		return nil, errors.New("user not found")
	}

	// Validasi password (hash dibandingkan)
	if user.Password != password {
		return nil, errors.New("invalid credentials")
	}

	return &user, nil
}
