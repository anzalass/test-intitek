package product

type ProductRequest struct {
	Name     string `json:"name" validate:"required"`
	SKU      string `json:"sku" validate:"required"`
	Quantity int    `json:"quantity" validate:"required,gte=0"`
	Location string `json:"location" validate:"required"`
	Status   string `json:"status" validate:"required,oneof=Available 'Out of Stock'"`
}
