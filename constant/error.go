package constant

import "fmt"

type ErrorMsg error

var (
	ErrorCounterNotFound   ErrorMsg = fmt.Errorf("counter not found")
	ErrorStockNotEnough    ErrorMsg = fmt.Errorf("stock not enough")
	ErrorUserNotFound      ErrorMsg = fmt.Errorf("user not found")
	ErrorProductNotFound   ErrorMsg = fmt.Errorf("product not found")
	ErrorProductsNotFound  ErrorMsg = fmt.Errorf("some product(s) not found")
	ErrorFailedCreateOrder ErrorMsg = fmt.Errorf("failed create order")
)
