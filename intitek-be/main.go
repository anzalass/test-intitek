package main

import (
	"fmt"
	"net/http"
	"strconv"

	"example.com/m/v2/module/auth"
	"example.com/m/v2/module/product"
	"example.com/m/v2/routes"
	"example.com/m/v2/utils"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
    utils.InitConfig()
    e := echo.New()
    db:=utils.ConnectDB()

    e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Test Intitek")
	})
	corsConfig := middleware.CORSConfig{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowCredentials: true,
	}
	e.Use(middleware.CORSWithConfig(corsConfig))


    productRepository := product.NewProductRepository(db)
	productService := product.NewProductService(productRepository)
	productHandler := product.NewProductHandler(productService)


    authRepository := auth.NewAuthRepository(db)
	authService := auth.NewAuthService(authRepository)
	authHandler := auth.NewAuthHandler(authService)

    routes.RouteProduct(e,productHandler)
    routes.RouteAuth(e,authHandler)



    port := strconv.Itoa(utils.InitConfig().ServerPort)
    utils.Migrate(db)
    e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", port)))


}