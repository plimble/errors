package errors

import (
	"fmt"
)

type Error interface {
	Status() int
	Error() string
	Code() string
}

type Errors struct {
	ErrStatus  int    `json:"-"`
	ErrMessage string `json:"message"`
	ErrCode    string `json:"code,omitempty"`
}

func (e *Errors) Status() int {
	return e.ErrStatus
}

func (e *Errors) Code() string {
	return e.ErrCode
}

func (e *Errors) Error() string {
	return e.ErrMessage
}

func New(message string) error {
	return &Errors{
		ErrStatus:  500,
		ErrMessage: message,
	}
}

func NewStatus(status int, message string) error {
	return &Errors{
		ErrStatus:  status,
		ErrMessage: message,
	}
}

func NewStatusf(status int, format string, v ...interface{}) error {
	return &Errors{
		ErrStatus:  status,
		ErrMessage: fmt.Sprintf(format, v...),
	}
}

func Newf(format string, v ...interface{}) error {
	return &Errors{
		ErrStatus:  500,
		ErrMessage: fmt.Sprintf(format, v...),
	}
}

func BadRequest(message string) error {
	return &Errors{
		ErrStatus:  400,
		ErrMessage: message,
	}
}

func Unauthorized(message string) error {
	return &Errors{
		ErrStatus:  401,
		ErrMessage: message,
	}
}

func Forbidden(message string) error {
	return &Errors{
		ErrStatus:  403,
		ErrMessage: message,
	}
}

func NotFound(message string) error {
	return &Errors{
		ErrStatus:  404,
		ErrMessage: message,
	}
}

func InternalError(message string) error {
	return &Errors{
		ErrStatus:  500,
		ErrMessage: message,
	}
}

func BadRequestf(format string, v ...interface{}) error {
	return &Errors{
		ErrStatus:  400,
		ErrMessage: fmt.Sprintf(format, v...),
	}
}

func Unauthorizedf(format string, v ...interface{}) error {
	return &Errors{
		ErrStatus:  401,
		ErrMessage: fmt.Sprintf(format, v...),
	}
}

func Forbiddenf(format string, v ...interface{}) error {
	return &Errors{
		ErrStatus:  403,
		ErrMessage: fmt.Sprintf(format, v...),
	}
}

func NotFoundf(format string, v ...interface{}) error {
	return &Errors{
		ErrStatus:  404,
		ErrMessage: fmt.Sprintf(format, v...),
	}
}

func InternalErrorf(format string, v ...interface{}) error {
	return &Errors{
		ErrStatus:  500,
		ErrMessage: fmt.Sprintf(format, v...),
	}
}

// With Code

func NewCode(code string, message string) error {
	return &Errors{
		ErrStatus:  500,
		ErrCode:    code,
		ErrMessage: message,
	}
}

func NewStatusCode(status int, code, message string) error {
	return &Errors{
		ErrStatus:  500,
		ErrCode:    code,
		ErrMessage: message,
	}
}

func NewStatusCodef(status int, code, format string, v ...interface{}) error {
	return &Errors{
		ErrStatus:  status,
		ErrCode:    code,
		ErrMessage: fmt.Sprintf(format, v...),
	}
}

func NewCodef(code, format string, v ...interface{}) error {
	return &Errors{
		ErrStatus:  500,
		ErrCode:    code,
		ErrMessage: fmt.Sprintf(format, v...),
	}
}

func BadRequestCode(code, message string) error {
	return &Errors{
		ErrStatus:  400,
		ErrCode:    code,
		ErrMessage: message,
	}
}

func UnauthorizedCode(code, message string) error {
	return &Errors{
		ErrStatus:  401,
		ErrCode:    code,
		ErrMessage: message,
	}
}

func ForbiddenCode(code, message string) error {
	return &Errors{
		ErrStatus:  403,
		ErrCode:    code,
		ErrMessage: message,
	}
}

func NotFoundCode(code, message string) error {
	return &Errors{
		ErrStatus:  404,
		ErrCode:    code,
		ErrMessage: message,
	}
}

func InternalErrorCode(code, message string) error {
	return &Errors{
		ErrStatus:  500,
		ErrCode:    code,
		ErrMessage: message,
	}
}

func BadRequestCodef(code, format string, v ...interface{}) error {
	return &Errors{
		ErrStatus:  400,
		ErrCode:    code,
		ErrMessage: fmt.Sprintf(format, v...),
	}
}

func UnauthorizedCodef(code, format string, v ...interface{}) error {
	return &Errors{
		ErrStatus:  401,
		ErrCode:    code,
		ErrMessage: fmt.Sprintf(format, v...),
	}
}

func ForbiddenCodef(code, format string, v ...interface{}) error {
	return &Errors{
		ErrStatus:  403,
		ErrCode:    code,
		ErrMessage: fmt.Sprintf(format, v...),
	}
}

func NotFoundCodef(code, format string, v ...interface{}) error {
	return &Errors{
		ErrStatus:  404,
		ErrCode:    code,
		ErrMessage: fmt.Sprintf(format, v...),
	}
}

func InternalErrorCodef(code, format string, v ...interface{}) error {
	return &Errors{
		ErrStatus:  500,
		ErrCode:    code,
		ErrMessage: fmt.Sprintf(format, v...),
	}
}

func IsNotFound(err error) bool {
	return checkStatus(404, err)
}

func IsBadRequest(err error) bool {
	return checkStatus(404, err)
}

func IsInternalError(err error) bool {
	return checkStatus(404, err)
}

func checkStatus(status int, err error) bool {
	errs, ok := err.(Error)
	if ok {
		if errs.Status() == status {
			return true
		}
	}

	return false
}
