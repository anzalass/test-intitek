package product

import (
	"net/http"
	"strconv"

	"example.com/m/v2/models"
	"example.com/m/v2/utils"
	"github.com/labstack/echo/v4"
)

type ProductHandler struct {
	service ServiceProductInterface
}

func NewProductHandler(service ServiceProductInterface) HandlerProductInterface {
	return &ProductHandler{
		service: service,
	}
}

// CreateProduct - Menambahkan produk baru ke inventory
func (h *ProductHandler) CreateProduct() echo.HandlerFunc {
	return func(c echo.Context) error {
		var req ProductRequest
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

		product := &models.ProductModel{
			Name:     req.Name,
			SKU:      req.SKU,
			Quantity: req.Quantity,
			Location: req.Location,
			Status:   req.Status,
		}

		res, err := h.service.CreateProduct(product)
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

// GetProducts - Mengambil daftar produk dengan filter opsional
func (h *ProductHandler) GetProducts() echo.HandlerFunc {
	return func(c echo.Context) error {
		filterStatus := c.QueryParam("status")
		lowStock := c.QueryParam("low_stock") == "true"
		
		// Ambil page dan pageSize, default ke 1 dan 10 jika tidak ada
		page, _ := strconv.Atoi(c.QueryParam("page"))
		if page < 1 {
			page = 1
		}
		pageSize, _ := strconv.Atoi(c.QueryParam("pageSize"))
		if pageSize < 1 {
			pageSize = 10
		}

		products, total, err := h.service.GetProducts(filterStatus, lowStock, page, pageSize)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]any{
				"success": false,
				"message": "Gagal mengambil produk",
				"error":   err.Error(),
			})
		}

		return c.JSON(http.StatusOK, map[string]any{
			"success": true,
			"data":    products,
			"page":    page,
			"pageSize": pageSize,
			"total":   total,
		})
	}
}


// GetProductBySKU - Mengambil satu produk berdasarkan SKU
func (h *ProductHandler) GetProductBySKU() echo.HandlerFunc {
	return func(c echo.Context) error {
		sku := c.Param("sku")
		if sku == "" {
			return c.JSON(http.StatusBadRequest, map[string]any{
				"error": "SKU is required",
			})
		}

		res, err := h.service.GetProductBySKU(sku)
		if err != nil {
			return c.JSON(http.StatusNotFound, map[string]any{
				"error": "Product not found",
			})
		}

		return c.JSON(http.StatusOK, map[string]any{
			"success": true,
			"data":    res,
		})
	}
}

// UpdateProduct - Memperbarui detail produk berdasarkan SKU
func (h *ProductHandler) UpdateProduct() echo.HandlerFunc {
	return func(c echo.Context) error {
		sku := c.Param("sku")
		if sku == "" {
			return c.JSON(http.StatusBadRequest, map[string]any{
				"error": "SKU is required",
			})
		}

		var req ProductRequest
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

		product := &models.ProductModel{
			Name:     req.Name,
			SKU:      req.SKU,
			Quantity: req.Quantity,
			Location: req.Location,
			Status:   req.Status,
		}

		res, err := h.service.UpdateProduct(sku, product)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]any{
				"error": err.Error(),
			})
		}

		return c.JSON(http.StatusOK, map[string]any{
			"success": true,
			"data":    res,
		})
	}
}

// DeleteProduct - Menghapus produk berdasarkan SKU
func (h *ProductHandler) DeleteProduct() echo.HandlerFunc {
	return func(c echo.Context) error {
		sku := c.Param("sku")
		if sku == "" {
			return c.JSON(http.StatusBadRequest, map[string]any{
				"error": "SKU is required",
			})
		}

		err := h.service.DeleteProduct(sku)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]any{
				"error": err.Error(),
			})
		}

		return c.JSON(http.StatusOK, map[string]any{
			"success": true,
			"message": "Product successfully deleted",
		})
	}
}
