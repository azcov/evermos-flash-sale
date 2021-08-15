package util

import (
	"fmt"
	"net/http"
	"time"

	"github.com/azcov/evermos-flash-sale/constant"
	"github.com/azcov/evermos-flash-sale/domain/response"
	"github.com/labstack/echo/v4"
	log "go.uber.org/zap"
)

//ParsingError is
type ParsingError struct {
	msg string
}

func (re *ParsingError) Error() string { return re.msg }

// SuccessResponse returns
func SuccessResponse(ctx echo.Context, message string, data interface{}) error {

	responseData := response.Base{
		Status:     constant.StatusSuccessText,
		StatusCode: http.StatusOK,
		Message:    message,
		Timestamp:  time.Now().UTC(),
		Data:       data,
	}

	log.S().Info("success response")

	return ctx.JSON(http.StatusOK, responseData)
}

// CreatedResponse returns
func CreatedResponse(ctx echo.Context, message string, data interface{}) error {

	responseData := response.Base{
		Status:     constant.StatusCreatedText,
		StatusCode: http.StatusCreated,
		Message:    message,
		Timestamp:  time.Now().UTC(),
		Data:       data,
	}

	log.S().Info("success create data")

	return ctx.JSON(http.StatusCreated, responseData)
}

// ErrorResponse returns
func ErrorResponse(ctx echo.Context, err error, data interface{}) error {
	statusCode, err := errorType(err)
	switch statusCode {
	case http.StatusConflict:
		return ErrorConflictResponse(ctx, err, data)
	case http.StatusBadRequest:
		return ErrorBadRequest(ctx, err, data)
	case http.StatusNotFound:
		return ErrorNotFound(ctx, err, data)
	case http.StatusUnauthorized:
		return ErrorUnauthorized(ctx, err, data)
	case http.StatusForbidden:
		return ErrorForbidden(ctx, err, data)
	}
	return ErrorInternalServerResponse(ctx, err, data)
}

// ErrorDefaultResponse returns
func ErrorDefaultResponse(ctx echo.Context, statusCode int, message string) error {

	switch statusCode {
	case http.StatusConflict:
		return ErrorConflictResponse(ctx, fmt.Errorf(message), nil)
	case http.StatusBadRequest:
		return ErrorBadRequest(ctx, fmt.Errorf(message), nil)
	case http.StatusNotFound:
		return ErrorNotFound(ctx, fmt.Errorf(message), nil)
	case http.StatusUnauthorized:
		return ErrorUnauthorized(ctx, fmt.Errorf(message), nil)
	}
	return ErrorInternalServerResponse(ctx, fmt.Errorf(message), nil)
}

// ErrorConflictResponse returns
func ErrorConflictResponse(ctx echo.Context, err error, data interface{}) error {

	responseData := response.Base{
		Status:     constant.StatusConflictText,
		StatusCode: http.StatusConflict,
		Message:    err.Error(),
		Timestamp:  time.Now().UTC(),
		Data:       data,
	}

	log.S().Errorf("conflict data error : %s ", err.Error())

	return ctx.JSON(http.StatusConflict, responseData)
}

// ErrorInternalServerResponse returns
func ErrorInternalServerResponse(ctx echo.Context, err error, data interface{}) error {

	responseData := response.Base{
		Status:     constant.StatusInternalServerErrorText,
		StatusCode: http.StatusInternalServerError,
		Message:    err.Error(),
		Timestamp:  time.Now().UTC(),
		Data:       data,
	}

	log.S().Errorf("internal server error : %s ", err.Error())

	return ctx.JSON(http.StatusInternalServerError, responseData)
}

// ErrorBadRequest returns
func ErrorBadRequest(ctx echo.Context, err error, data interface{}) error {
	responseData := response.Base{
		Status:     constant.StatusBadRequestText,
		StatusCode: http.StatusBadRequest,
		Message:    err.Error(),
		Timestamp:  time.Now().UTC(),
		Data:       data,
	}

	log.S().Errorf("bad request error : %s ", err.Error())

	return ctx.JSON(http.StatusBadRequest, responseData)
}

// ErrorNotFound returns
func ErrorNotFound(ctx echo.Context, err error, data interface{}) error {
	responseData := response.Base{
		Status:     constant.StatusNotFoundText,
		StatusCode: http.StatusNotFound,
		Message:    err.Error(),
		Timestamp:  time.Now().UTC(),
		Data:       data,
	}

	log.S().Errorf("error not found : %s ", err.Error())

	return ctx.JSON(http.StatusNotFound, responseData)
}

// ErrorParsing returns
func ErrorParsing(ctx echo.Context, err error, data interface{}) error {

	responseData := response.Base{
		Status:     constant.StatusUnprocessableEntityText,
		StatusCode: http.StatusUnprocessableEntity,
		Message:    err.Error(),
		Timestamp:  time.Now().UTC(),
		Data:       data,
	}

	log.S().Errorf("parsing data error : %s ", err.Error())

	return ctx.JSON(http.StatusUnprocessableEntity, responseData)
}

// ErrorValidate returns
func ErrorValidate(ctx echo.Context, err error, data interface{}) error {
	message := switchErrorValidation(err)
	responseData := response.Base{
		Status:     constant.StatusBadRequestText,
		StatusCode: http.StatusBadRequest,
		Message:    message,
		Timestamp:  time.Now().UTC(),
		Data:       data,
	}

	log.S().Errorf("validate data error : %s ", err.Error())

	return ctx.JSON(http.StatusBadRequest, responseData)
}

// ErrorUnauthorized returns
func ErrorUnauthorized(ctx echo.Context, err error, data interface{}) error {
	responseData := response.Base{
		Status:     constant.StatusUnauthorized,
		StatusCode: http.StatusUnauthorized,
		Message:    err.Error(),
		Timestamp:  time.Now().UTC(),
		Data:       data,
	}

	log.S().Errorf("unauthorized error : %s ", err.Error())

	return ctx.JSON(http.StatusUnauthorized, responseData)
}

// ErrorForbidden returns
func ErrorForbidden(ctx echo.Context, err error, data interface{}) error {
	responseData := response.Base{
		Status:     constant.StatusForbidden,
		StatusCode: http.StatusForbidden,
		Message:    err.Error(),
		Timestamp:  time.Now().UTC(),
		Data:       data,
	}

	log.S().Errorf("forbidden error : %s ", err.Error())

	return ctx.JSON(http.StatusForbidden, responseData)
}

// SetCookie returns
func SetCookie(ctx echo.Context, data map[string]string) {
	for key, value := range data {
		cookie := &http.Cookie{}
		cookie.Name = key
		cookie.Value = value
		cookie.HttpOnly = true
		cookie.Path = "/"
		ctx.SetCookie(cookie)
	}
}
