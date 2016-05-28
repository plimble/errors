package errors

import (
	"fmt"

	"github.com/Sirupsen/logrus"
)

type Error interface {
	Status() int
	Error() string
	Code() string
	NotFound() bool
	BadRequest() bool
	Unauthorized() bool
	Forbidden() bool
	Logrus(err error) *logrus.Entry
}

type Errors struct {
	ErrStatus  int    `json:"-"`
	ErrMessage string `json:"message"`
	ErrCode    string `json:"code,omitempty"`
}

func (e *Errors) Logrus(err error) *logrus.Entry {
	return logrus.WithField("code", e.Code()).WithError(err)
}

func (e *Errors) NotFound() bool {
	if e.ErrStatus == 404 {
		return true
	}

	return false
}

func (e *Errors) BadRequest() bool {
	if e.ErrStatus == 400 {
		return true
	}

	return false
}

func (e *Errors) Unauthorized() bool {
	if e.ErrStatus == 401 {
		return true
	}

	return false
}

func (e *Errors) Forbidden() bool {
	if e.ErrStatus == 403 {
		return true
	}

	return false
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

func New(message string) Error {
	return &Errors{
		ErrStatus:  500,
		ErrMessage: message,
	}
}

func NewStatus(status int, message string) Error {
	return &Errors{
		ErrStatus:  status,
		ErrMessage: message,
	}
}

func NewStatusf(status int, format string, v ...interface{}) Error {
	return &Errors{
		ErrStatus:  status,
		ErrMessage: fmt.Sprintf(format, v...),
	}
}

func Newf(format string, v ...interface{}) Error {
	return &Errors{
		ErrStatus:  500,
		ErrMessage: fmt.Sprintf(format, v...),
	}
}

func BadRequest(message string) Error {
	return &Errors{
		ErrStatus:  400,
		ErrMessage: message,
	}
}

func Unauthorized(message string) Error {
	return &Errors{
		ErrStatus:  401,
		ErrMessage: message,
	}
}

func Forbidden(message string) Error {
	return &Errors{
		ErrStatus:  403,
		ErrMessage: message,
	}
}

func NotFound(message string) Error {
	return &Errors{
		ErrStatus:  404,
		ErrMessage: message,
	}
}

func InternalError(message string) Error {
	return &Errors{
		ErrStatus:  500,
		ErrMessage: message,
	}
}

func BadRequestf(format string, v ...interface{}) Error {
	return &Errors{
		ErrStatus:  400,
		ErrMessage: fmt.Sprintf(format, v...),
	}
}

func Unauthorizedf(format string, v ...interface{}) Error {
	return &Errors{
		ErrStatus:  401,
		ErrMessage: fmt.Sprintf(format, v...),
	}
}

func Forbiddenf(format string, v ...interface{}) Error {
	return &Errors{
		ErrStatus:  403,
		ErrMessage: fmt.Sprintf(format, v...),
	}
}

func NotFoundf(format string, v ...interface{}) Error {
	return &Errors{
		ErrStatus:  404,
		ErrMessage: fmt.Sprintf(format, v...),
	}
}

func InternalErrorf(format string, v ...interface{}) Error {
	return &Errors{
		ErrStatus:  500,
		ErrMessage: fmt.Sprintf(format, v...),
	}
}

// With Code

func NewCode(code string, message string) Error {
	return &Errors{
		ErrStatus:  500,
		ErrCode:    code,
		ErrMessage: message,
	}
}

func NewStatusCode(status int, code, message string) Error {
	return &Errors{
		ErrStatus:  500,
		ErrCode:    code,
		ErrMessage: message,
	}
}

func NewStatusCodef(status int, code, format string, v ...interface{}) Error {
	return &Errors{
		ErrStatus:  status,
		ErrCode:    code,
		ErrMessage: fmt.Sprintf(format, v...),
	}
}

func NewCodef(code, format string, v ...interface{}) Error {
	return &Errors{
		ErrStatus:  500,
		ErrCode:    code,
		ErrMessage: fmt.Sprintf(format, v...),
	}
}

func BadRequestCode(code, message string) Error {
	return &Errors{
		ErrStatus:  400,
		ErrCode:    code,
		ErrMessage: message,
	}
}

func UnauthorizedCode(code, message string) Error {
	return &Errors{
		ErrStatus:  401,
		ErrCode:    code,
		ErrMessage: message,
	}
}

func ForbiddenCode(code, message string) Error {
	return &Errors{
		ErrStatus:  403,
		ErrCode:    code,
		ErrMessage: message,
	}
}

func NotFoundCode(code, message string) Error {
	return &Errors{
		ErrStatus:  404,
		ErrCode:    code,
		ErrMessage: message,
	}
}

func InternalErrorCode(code, message string) Error {
	return &Errors{
		ErrStatus:  500,
		ErrCode:    code,
		ErrMessage: message,
	}
}

func BadRequestCodef(code, format string, v ...interface{}) Error {
	return &Errors{
		ErrStatus:  400,
		ErrCode:    code,
		ErrMessage: fmt.Sprintf(format, v...),
	}
}

func UnauthorizedCodef(code, format string, v ...interface{}) Error {
	return &Errors{
		ErrStatus:  401,
		ErrCode:    code,
		ErrMessage: fmt.Sprintf(format, v...),
	}
}

func ForbiddenCodef(code, format string, v ...interface{}) Error {
	return &Errors{
		ErrStatus:  403,
		ErrCode:    code,
		ErrMessage: fmt.Sprintf(format, v...),
	}
}

func NotFoundCodef(code, format string, v ...interface{}) Error {
	return &Errors{
		ErrStatus:  404,
		ErrCode:    code,
		ErrMessage: fmt.Sprintf(format, v...),
	}
}

func InternalErrorCodef(code, format string, v ...interface{}) Error {
	return &Errors{
		ErrStatus:  500,
		ErrCode:    code,
		ErrMessage: fmt.Sprintf(format, v...),
	}
}

func IsNotFound(err error) bool {
	if cerr, ok := err.(Error); ok && cerr.Status() == 404 {
		return true
	}

	return false
}
