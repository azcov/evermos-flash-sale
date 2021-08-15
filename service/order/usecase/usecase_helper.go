package usecase

import (
	"github.com/azcov/evermos-flash-sale/constant"
	"github.com/azcov/evermos-flash-sale/domain/entity"
	"github.com/azcov/evermos-flash-sale/domain/request"
)

func (u *usecase) getProductIdsFromRequest(req *request.CreateOrderRequest) []int32 {
	var res []int32
	for i := range req.Products {
		res = append(res, req.Products[i].ProductID)
	}
	return res
}

func (u *usecase) checkAndCountProducts(req *request.CreateOrderRequest, products map[int32]*entity.GetProductByProductIDForUpdateRow) (totalPrice int64, productIds, qtys []int32, prices []int64, err error) {
	for i := range req.Products {
		if products[req.Products[i].ProductID].Qty < req.Products[i].Qty {
			err = constant.ErrorStockNotEnough
			return
		}
		productIds = append(productIds, products[req.Products[i].ProductID].ID)
		qtys = append(qtys, products[req.Products[i].ProductID].Qty)
		prices = append(prices, products[req.Products[i].ProductID].Price)
		totalPrice += int64(products[req.Products[i].ProductID].Qty) * products[req.Products[i].ProductID].Price
	}
	return
}
