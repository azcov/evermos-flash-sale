package entity

type OrderProduct struct {
	OrderID   int32 `json:"order_id"`
	ProductID int32 `json:"product_id"`
	Price     int64 `json:"price"`
	Qty       int32 `json:"qty"`
	CreatedAt int64 `json:"created_at"`
}

type CreateOrderProductsParams struct {
	OrderID    int32   `json:"order_id"`
	ProductIds []int32 `json:"product_ids"`
	Prices     []int64 `json:"prices"`
	Qtys       []int32 `json:"qtys"`
	CreatedAt  int64   `json:"created_at"`
}
