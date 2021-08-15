package usecase

import (
	"context"

	request "github.com/azcov/evermos-flash-sale/domain/request"
)

type Usecase interface {
	CreateOrder(ctx context.Context, req *request.CreateOrderRequest) error
}
