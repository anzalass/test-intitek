package auth



type RegisterRequest struct {
	Username string `json:"username" validate:"required,min=3,max=255"` 
	Password string `json:"password" validate:"required,min=6"`        
}

type LoginRequest struct {
	Username string `json:"username" validate:"required"` 
	Password string `json:"password" validate:"required"` 
}
