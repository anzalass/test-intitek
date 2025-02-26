package product

import (
	"example.com/m/v2/models"
	"github.com/labstack/echo/v4"
)



type RepositoryProductInterface interface {
	CreateProduct(newProduct *models.ProductModel) (*models.ProductModel, error)
	 GetProducts(filterStatus string, lowStock bool, page, pageSize int) ([]models.ProductModel, int64, error) 
	GetProductBySKU(sku string) (*models.ProductModel, error)
	UpdateProduct(sku string, updatedProduct *models.ProductModel) (*models.ProductModel, error)
	DeleteProduct(sku string) error
}



type ServiceProductInterface interface {
	CreateProduct(newProduct *models.ProductModel) (*models.ProductModel, error)
	 GetProducts(filterStatus string, lowStock bool, page, pageSize int) ([]models.ProductModel, int64, error) 
	GetProductBySKU(sku string) (*models.ProductModel, error)
	UpdateProduct(sku string, updatedProduct *models.ProductModel) (*models.ProductModel, error)
	DeleteProduct(sku string) error
}


type HandlerProductInterface interface {
	CreateProduct() echo.HandlerFunc
	GetProducts() echo.HandlerFunc
	GetProductBySKU() echo.HandlerFunc
	UpdateProduct() echo.HandlerFunc
	DeleteProduct() echo.HandlerFunc
}
