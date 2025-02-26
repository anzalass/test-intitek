package auth

import (
	"net/http"

	"example.com/m/v2/models"
	"example.com/m/v2/utils"
	"github.com/labstack/echo/v4"
)

type AuthHandler struct {
	service ServiceAuthInterface
}

func NewAuthHandler(service ServiceAuthInterface) HandlerAuthInterface {
	return &AuthHandler{service: service}
}

// RegisterUser - Handler untuk registrasi user
func (h *AuthHandler) RegisterUser() echo.HandlerFunc {
	return func(c echo.Context) error {
		var req RegisterRequest
		if err := c.Bind(&req); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]any{
				"error": "Invalid request format",
			})
		}

		// Validasi request
		if err := utils.ValidateStruct(req); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]any{
				"error": err.Error(),
			})
		}

		newUser := &models.UserModel{
			Username: req.Username,
			Password: req.Password,
		}

		res, err := h.service.RegisterUser(newUser)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]any{
				"error": err.Error(),
			})
		}

		return c.JSON(http.StatusCreated, map[string]any{
			"success": true,
			"data":    res,
		})
	}
}

func (h *AuthHandler) LoginUser() echo.HandlerFunc {
	return func(c echo.Context) error {
		var req LoginRequest
		if err := c.Bind(&req); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]any{
				"error": "Invalid request format",
			})
		}

		// Validasi request
		if err := utils.ValidateStruct(req); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]any{
				"error": err.Error(),
			})
		}

		user, err := h.service.LoginUser(req.Username, req.Password)
		if err != nil {
			return c.JSON(http.StatusUnauthorized, map[string]any{
				"error": err.Error(),
			})
		}

		return c.JSON(http.StatusOK, map[string]any{
			"success": true,
			"user":    user,
		})
	}
}
