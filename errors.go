package errors

import (
	"fmt"
)

type Error interface {
	Code() int
	Error() string
}

type Errors struct {
	ErrCode    int    `json:"-"`
	ErrMessage string `json:"message"`
}

func (e *Errors) Code() int {
	return e.ErrCode
}

func (e *Errors) Error() string {
	return e.ErrMessage
}

func New(message string) error {
	return &Errors{
		ErrCode:    500,
		ErrMessage: message,
	}
}

func NewCode(code int, message string) error {
	return &Errors{
		ErrCode:    500,
		ErrMessage: message,
	}
}

func NewCodef(code int, format string, v ...interface{}) error {
	return &Errors{
		ErrCode:    code,
		ErrMessage: fmt.Sprintf(format, v...),
	}
}

func Newf(format string, v ...interface{}) error {
	return &Errors{
		ErrCode:    500,
		ErrMessage: fmt.Sprintf(format, v...),
	}
}

func BadRequest(message string) error {
	return &Errors{
		ErrCode:    400,
		ErrMessage: message,
	}
}

func Unauthorized(message string) error {
	return &Errors{
		ErrCode:    401,
		ErrMessage: message,
	}
}

func Forbidden(message string) error {
	return &Errors{
		ErrCode:    403,
		ErrMessage: message,
	}
}

func NotFound(message string) error {
	return &Errors{
		ErrCode:    404,
		ErrMessage: message,
	}
}

func InternalError(message string) error {
	return &Errors{
		ErrCode:    500,
		ErrMessage: message,
	}
}

func BadRequestf(format string, v ...interface{}) error {
	return &Errors{
		ErrCode:    400,
		ErrMessage: fmt.Sprintf(format, v...),
	}
}

func Unauthorizedf(format string, v ...interface{}) error {
	return &Errors{
		ErrCode:    401,
		ErrMessage: fmt.Sprintf(format, v...),
	}
}

func Forbiddenf(format string, v ...interface{}) error {
	return &Errors{
		ErrCode:    403,
		ErrMessage: fmt.Sprintf(format, v...),
	}
}

func NotFoundf(format string, v ...interface{}) error {
	return &Errors{
		ErrCode:    404,
		ErrMessage: fmt.Sprintf(format, v...),
	}
}

func InternalErrorf(format string, v ...interface{}) error {
	return &Errors{
		ErrCode:    500,
		ErrMessage: fmt.Sprintf(format, v...),
	}
}
