package models


type ProductModel struct {
    Name     string `gorm:"type:varchar(255);not null" json:"name"`  // Nama produk
    SKU      string `gorm:"type:varchar(100);unique;not null" json:"sku"` // SKU harus unik
    Quantity int    `gorm:"not null" json:"quantity"` // Jumlah stok produk
    Location string `gorm:"type:varchar(255);not null" json:"location"` // Lokasi penyimpanan
    Status   string `gorm:"type:varchar(50);not null" json:"status"`  // Status produk (misalnya: "Available", "Out of Stock")
}

type UserModel struct {
	ID       uint   `gorm:"primary_key"`
	Username string `gorm:"type:varchar(255);unique;not null"`
	Password string `gorm:"type:varchar(255);not null"`
}



func (ProductModel) TableName() string {
	return "products"
}

func (UserModel) TableName() string {
	return "users"
}