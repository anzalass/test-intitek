package auth

import (
	"example.com/m/v2/models"
)

type AuthService struct {
	repo RepositoryAuthInterface
}

func NewAuthService(repo RepositoryAuthInterface) ServiceAuthInterface {
	return &AuthService{repo: repo}
}

// RegisterUser - Menambahkan user baru dengan enkripsi password
func (s *AuthService) RegisterUser(newUser *models.UserModel) (*models.UserModel, error) {


	// Simpan user ke database
	user, err := s.repo.RegisterUser(newUser)
	if err != nil {
		return nil, err
	}

	return user, nil
}

// LoginUser - Melakukan autentikasi user
func (s *AuthService) LoginUser(username, password string) (*models.UserModel, error) {
	user, err := s.repo.LoginUser(username, password)
	if err != nil {
		return nil, err
	}

	return user, nil
}
