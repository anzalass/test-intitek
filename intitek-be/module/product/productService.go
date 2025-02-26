package product

import (
	"errors"

	"example.com/m/v2/models"
)

type ProductService struct {
	repo RepositoryProductInterface
}

func NewProductService(repo RepositoryProductInterface) ServiceProductInterface {
	return &ProductService{
		repo: repo,
	}
}

// CreateProduct menambahkan produk baru ke database
func (s *ProductService) CreateProduct(newProduct *models.ProductModel) (*models.ProductModel, error) {
	// Cek apakah SKU sudah ada
	existingProduct, _ := s.repo.GetProductBySKU(newProduct.SKU)
	if existingProduct != nil {
		return nil, errors.New("produk dengan SKU ini sudah ada")
	}

	// Simpan ke database
	res, err := s.repo.CreateProduct(newProduct)
	if err != nil {
		return nil, errors.New("gagal menambahkan produk")
	}

	return res, nil
}

// GetProducts mengambil daftar produk dengan filter status atau stok rendah
func (s *ProductService) GetProducts(filterStatus string, lowStock bool, page, pageSize int) ([]models.ProductModel, int64, error)  {
	return s.repo.GetProducts(filterStatus, lowStock, page, pageSize)
}

// GetProductBySKU mengambil satu produk berdasarkan SKU
func (s *ProductService) GetProductBySKU(sku string) (*models.ProductModel, error) {
	res, err := s.repo.GetProductBySKU(sku)
	if err != nil {
		return nil, errors.New("produk tidak ditemukan")
	}
	return res, nil
}

// UpdateProduct memperbarui data produk berdasarkan SKU
func (s *ProductService) UpdateProduct(sku string, updatedProduct *models.ProductModel) (*models.ProductModel, error) {
	// Cari produk berdasarkan SKU
	existingProduct, err := s.repo.GetProductBySKU(sku)
	if err != nil {
		return nil, errors.New("produk tidak ditemukan")
	}

	// Update data (hanya field yang berubah)
	s.repo.UpdateProduct(sku, updatedProduct)

	return existingProduct, nil
}

// DeleteProduct menghapus produk berdasarkan SKU
func (s *ProductService) DeleteProduct(sku string) error {
	// Cek apakah produk ada
	_, err := s.repo.GetProductBySKU(sku)
	if err != nil {
		return errors.New("produk tidak ditemukan")
	}

	// Hapus produk
	return s.repo.DeleteProduct(sku)
}
