package product

import (
	"example.com/m/v2/models"
	"gorm.io/gorm"
)

type ProductRepository struct {
	db *gorm.DB
}

// NewProductRepository membuat instance baru dari ProductRepository
func NewProductRepository(db *gorm.DB) RepositoryProductInterface {
	return &ProductRepository{
		db: db,
	}
}

// CreateProduct menambahkan produk baru ke database
func (r *ProductRepository) CreateProduct(newProduct *models.ProductModel) (*models.ProductModel, error) {
	if err := r.db.Create(newProduct).Error; err != nil {
		return nil, err
	}
	return newProduct, nil
}

// GetProducts mengambil daftar produk dengan filter status atau stok rendah
func (r *ProductRepository) GetProducts(filterStatus string, lowStock bool, page, pageSize int) ([]models.ProductModel, int64, error) {
	var products []models.ProductModel
	var total int64

	query := r.db.Model(&models.ProductModel{}) // Tambahkan Model agar total count bisa dihitung

	// Filter berdasarkan status
	if filterStatus != "" {
		query = query.Where("status = ?", filterStatus)
	}

	// Filter berdasarkan stok rendah
	if lowStock {
		query = query.Where("quantity < ?", 10) // Contoh: stok dianggap rendah jika < 10
	}

	// Hitung total data tanpa pagination
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// Pagination logic
	offset := (page - 1) * pageSize
	if err := query.Offset(offset).Limit(pageSize).Find(&products).Error; err != nil {
		return nil, 0, err
	}

	return products, total, nil
}


// GetProductBysku mengambil satu produk berdasarkan sku
func (r *ProductRepository) GetProductBySKU(sku string) (*models.ProductModel, error) {
	var product models.ProductModel
	if err := r.db.Where("sku = ?", sku).First(&product).Error; err != nil {
		return nil, err
	}
	return &product, nil
}

// UpdateProduct memperbarui data produk berdasarkan sku
func (r *ProductRepository) UpdateProduct(sku string, updatedProduct *models.ProductModel) (*models.ProductModel, error) {
	var product models.ProductModel
	if err := r.db.Where("sku = ?", sku).First(&product).Error; err != nil {
		return nil, err
	}

	if err := r.db.Model(&models.ProductModel{}).Where("sku = ?", sku).Updates(updatedProduct).Error; err != nil {
	return nil, err
}


	return &product, nil
}

// DeleteProduct menghapus produk berdasarkan sku
func (r *ProductRepository) DeleteProduct(sku string) error {
	if err := r.db.Where("sku = ?", sku).Delete(&models.ProductModel{}).Error; err != nil {
		return err
	}
	return nil
}
