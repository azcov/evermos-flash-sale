package util

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/azcov/evermos-flash-sale/pkg/helper"
	"github.com/go-playground/validator/v10"
	"github.com/lib/pq"
)

func calculateMin(fieldName, num string) string {
	fieldsName := map[string]int{
		"phone number": 2,
	}

	i, err := strconv.Atoi(num)
	if err != nil {
		return num
	}

	for key := range fieldsName {
		if key == fieldName {
			return strconv.Itoa(i - fieldsName[key])
		}
	}
	return num
}

func oneofErrorSplit(choices string) (result string) {
	s := strings.Split(choices, " ")
	for idx := range s {
		message := fmt.Sprintf("%s or ", s[idx])
		if (idx + 1) == len(s) {
			message = s[idx]
		}
		result += message
	}
	return
}

func errorType(err error) (int, error) {
	switch {
	case isPqError(err):
		return helper.PqError(err)
	}
	return helper.CommonError(err)
}

func isPqError(err error) bool {
	if _, ok := err.(*pq.Error); ok {
		return true
	}
	return false
}

func switchErrorValidation(err error) (message string) {
	if castedObject, ok := err.(validator.ValidationErrors); ok {
		for _, err := range castedObject {
			field := err.Field()

			// Check Error Type
			switch err.Tag() {
			case "required":
				message = field + " is mandatory"
			default:
				message = err.Error()
			}
		}
	}
	return
}
