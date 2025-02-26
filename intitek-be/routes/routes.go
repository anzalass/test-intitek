package routes

import (
	"example.com/m/v2/module/auth"
	"example.com/m/v2/module/product"
	"github.com/labstack/echo/v4"
)

func RouteProduct(e *echo.Echo, h product.HandlerProductInterface) {
	e.POST("/product", h.CreateProduct())
	e.PUT("/product/:sku", h.UpdateProduct())
	e.GET("/product/:sku", h.GetProductBySKU())
	e.GET("/products", h.GetProducts())
	e.DELETE("/product/:sku", h.DeleteProduct())
}

func RouteAuth(e *echo.Echo, h auth.HandlerAuthInterface) {
	e.POST("/register", h.RegisterUser())  // Endpoint untuk registrasi user
	e.POST("/login", h.LoginUser())        // Endpoint untuk login user
}