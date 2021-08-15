package usecase

import (
	"context"
	"time"

	"github.com/azcov/evermos-flash-sale/constant"
	appInit "github.com/azcov/evermos-flash-sale/init"
	"github.com/azcov/evermos-flash-sale/pkg/helper"
	"github.com/azcov/evermos-flash-sale/service/order/repository"

	"github.com/azcov/evermos-flash-sale/domain/entity"
	request "github.com/azcov/evermos-flash-sale/domain/request"
)

type usecase struct {
	orderRepo repository.Repository
	config    *appInit.Config
}

func NewOrderUsecase(config *appInit.Config, orderRepo repository.Repository) Usecase {
	return &usecase{
		orderRepo: orderRepo,
		config:    config,
	}
}

// --- CreateOrder used to --- //
func (u *usecase) CreateOrder(ctx context.Context, req *request.CreateOrderRequest) (err error) {

	user, err := u.orderRepo.GetuserByEmail(ctx, req.UserEmail)
	if err != nil {
		return
	}

	// * transaction init
	tx, err := u.orderRepo.BeginTx(ctx)
	if err != nil {
		return
	}

	counter, err := u.orderRepo.WithTx(tx).GetCounterByCounterID(ctx, int32(constant.CounterIDOrder))
	if err != nil {
		u.orderRepo.RollbackTx(tx)
		return
	}

	generatedOrderNo := helper.GenerateOrderNo(counter)

	productIds := u.getProductIdsFromRequest(req)

	products, err := u.orderRepo.WithTx(tx).GetProductsByProductIDForUpdate(ctx, productIds)
	if err != nil {
		u.orderRepo.RollbackTx(tx)
		return
	}
	total, productIds, qtys, prices, err := u.checkAndCountProducts(req, products)
	if err != nil {
		u.orderRepo.RollbackTx(tx)
		return
	}
	timeNow := time.Now()
	createOrderParams := &entity.CreateOrderParams{
		OrderNo:    generatedOrderNo,
		UserID:     user.ID,
		TotalPrice: total,
		CreatedAt:  timeNow.Unix(),
	}
	orderID, err := u.orderRepo.WithTx(tx).CreateOrder(ctx, createOrderParams)
	if err != nil {
		u.orderRepo.RollbackTx(tx)
		err = constant.ErrorFailedCreateOrder
		return
	}
	createOrderProductParams := &entity.CreateOrderProductsParams{
		OrderID:    orderID,
		ProductIds: productIds,
		Prices:     prices,
		Qtys:       qtys,
		CreatedAt:  timeNow.Unix(),
	}
	err = u.orderRepo.WithTx(tx).CreateOrderProducts(ctx, createOrderProductParams)
	if err != nil {
		u.orderRepo.RollbackTx(tx)
		err = constant.ErrorFailedCreateOrder
		return
	}
	for i := range req.Products {
		updateProductQtyParams := &entity.UpdateProductQtyParams{
			Qty:       products[req.Products[i].ProductID].Qty - req.Products[i].Qty,
			ProductID: req.Products[i].ProductID,
		}
		err = u.orderRepo.WithTx(tx).UpdateProductQty(ctx, updateProductQtyParams)
	}

	updateCounterParams := &entity.UpdateCounterByCounterIDParams{
		Counter:   counter.Counter + 1,
		UpdatedAt: timeNow.Unix(),
		CounterID: int32(constant.CounterIDOrder),
	}
	err = u.orderRepo.WithTx(tx).UpdateCounterByCounterID(ctx, updateCounterParams)
	if err != nil {
		u.orderRepo.RollbackTx(tx)
		err = constant.ErrorFailedCreateOrder
		return
	}

	u.orderRepo.CommitTx(tx)

	return
}
