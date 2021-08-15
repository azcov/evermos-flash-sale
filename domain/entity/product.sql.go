package entity

type Product struct {
	ID        int32  `json:"id"`
	Name      string `json:"name"`
	Price     int64  `json:"price"`
	Qty       int32  `json:"qty"`
	CreatedAt int64  `json:"created_at"`
}

type GetProductByProductIDForUpdateRow struct {
	ID    int32  `json:"id"`
	Name  string `json:"name"`
	Price int64  `json:"price"`
	Qty   int32  `json:"qty"`
}

type UpdateProductQtyParams struct {
	Qty       int32 `json:"qty"`
	ProductID int32 `json:"product_id"`
}
