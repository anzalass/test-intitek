package auth

import (
	"example.com/m/v2/models"
	"github.com/labstack/echo/v4"
)



type RepositoryAuthInterface interface {
	RegisterUser(newUser *models.UserModel) (*models.UserModel, error)
	LoginUser(username, password string) (*models.UserModel, error)
}



type ServiceAuthInterface interface {
	RegisterUser(newUser *models.UserModel) (*models.UserModel, error)
	LoginUser(username, password string) (*models.UserModel, error)
}


type HandlerAuthInterface interface {
	RegisterUser() echo.HandlerFunc
	LoginUser() echo.HandlerFunc
	
}
