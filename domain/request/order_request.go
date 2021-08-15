package request

type CreateOrderRequest struct {
	UserEmail string               `json:"user_email"`
	Products  []CreateOrderProduct `json:"products"`
}

type CreateOrderProduct struct {
	ProductID int32 `json:"product_id"`
	Qty       int32 `json:"qty"`
}
