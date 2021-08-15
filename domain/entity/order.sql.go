package entity

type Order struct {
	ID         int32  `json:"id"`
	OrderNo    string `json:"order_no"`
	UserID     string `json:"user_id"`
	TotalPrice int64  `json:"total_price"`
	CreatedAt  int64  `json:"created_at"`
}

type CreateOrderParams struct {
	OrderNo    string `json:"order_no"`
	UserID     int32  `json:"user_id"`
	TotalPrice int64  `json:"total_price"`
	CreatedAt  int64  `json:"created_at"`
}
