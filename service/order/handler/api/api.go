package http

import (
	"github.com/azcov/evermos-flash-sale/constant"
	"github.com/azcov/evermos-flash-sale/domain/request"
	"github.com/azcov/evermos-flash-sale/pkg/util"
	"github.com/azcov/evermos-flash-sale/service/order/usecase"
	"github.com/labstack/echo/v4"
)

// orderHandler  represent the httphandler for Order
type orderHandler struct {
	usecase usecase.Usecase
}

// NewOrderHandler will initialize the contact/ resources endpoint
func NewOrderHandler(e *echo.Group, us usecase.Usecase) {
	handler := &orderHandler{
		usecase: us,
	}
	routerV1 := e.Group("/v1")
	routerV1.POST("/checkout", handler.CreateOrder)

}

// CreateOrder godoc
// @Summary CreateOrder order
// @Description CreateOrder product order
// @Tags Order
// @Accept  json
// @Produce  json
// @Param request body request.CreateOrderRequest true "Request Body"
// @Success 201 {object} response.Base
// @Failure 400 {object} response.Base
// @Failure 409 {object} response.Base
// @Failure 422 {object} response.Base
// @Failure 500 {object} response.Base
// @Router /v1/checkout [post]
func (h *orderHandler) CreateOrder(c echo.Context) error {
	req := request.CreateOrderRequest{}
	ctx := c.Request().Context()

	// parsing
	err := util.ParsingParameter(c, &req)
	if err != nil {
		return util.ErrorParsing(c, err, nil)
	}

	// validate
	err = util.ValidateParameter(c, &req)
	if err != nil {
		return util.ErrorValidate(c, err, nil)
	}

	err = h.usecase.CreateOrder(ctx, &req)
	if err != nil {
		return util.ErrorResponse(c, err, nil)
	}

	return util.CreatedResponse(c, string(constant.SuccessCreateOrderResponse), nil)
}
