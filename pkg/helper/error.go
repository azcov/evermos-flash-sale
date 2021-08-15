package helper

import (
	"fmt"
	"net/http"

	"github.com/azcov/evermos-flash-sale/constant"
	"github.com/lib/pq"
)

var pqErrorMap = map[string]int{
	"unique_violation": http.StatusConflict,
}

// PqError is
func PqError(err error) (int, error) {
	if err, ok := err.(*pq.Error); ok {
		switch err.Code.Name() {
		case "unique_violation":
			return pqErrorMap["unique_violation"], fmt.Errorf("already exists")
		}
	}

	return http.StatusInternalServerError, fmt.Errorf(err.Error())
}

var commonErrorMap = map[error]int{
	constant.ErrorCounterNotFound:   http.StatusNotFound,
	constant.ErrorStockNotEnough:    http.StatusBadRequest,
	constant.ErrorUserNotFound:      http.StatusNotFound,
	constant.ErrorProductNotFound:   http.StatusNotFound,
	constant.ErrorProductsNotFound:  http.StatusNotFound,
	constant.ErrorFailedCreateOrder: http.StatusInternalServerError,
}

// CommonError is
func CommonError(err error) (int, error) {
	if commonErrorMap[err] != 0 {
		return commonErrorMap[err], fmt.Errorf(err.Error())
	}

	return http.StatusInternalServerError, fmt.Errorf(err.Error())
}
