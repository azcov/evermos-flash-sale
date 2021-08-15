package helper

import (
	"fmt"
	"time"

	"github.com/azcov/evermos-flash-sale/constant"
	"github.com/azcov/evermos-flash-sale/domain/entity"
)

func GenerateOrderNo(arg *entity.GetCounterByCounterIDRow) string {
	timeNow := time.Now()
	lastOrderTime := time.Unix(arg.UpdatedAt, 0)
	orderTime := lastOrderTime.Format(constant.DefaultOrderNoTimeFormat)
	if timeNow.Format(constant.DefaultOrderNoTimeFormat) != lastOrderTime.Format(constant.DefaultOrderNoTimeFormat) {
		orderTime = timeNow.Format(constant.DefaultOrderNoTimeFormat)
	}
	result := fmt.Sprintf(constant.DefaultOrderNoFormat, arg.Prefix, orderTime)
	orderNoStr := fmt.Sprintf("%d", arg.Counter)
	if len(result)+len(orderNoStr) > constant.DefaultOrderNoLength {
		return result + orderNoStr
	}

	if len(orderNoStr) == constant.DefaultOrderNoLength {
		return arg.Prefix + orderNoStr
	}
	m := constant.DefaultOrderNoLength - len(orderNoStr)
	for i := 0; i <= m; i++ {
		if i != m {
			result += "0"
			continue
		}
		result += orderNoStr
	}
	return result
}
