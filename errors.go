package errors

import (
	"errors"
	"fmt"
)

const (
	Internal     = 500
	BadReq       = 400
	NotFound     = 404
	Unauthorized = 401
	Forbidden    = 403
)

func New(s string) error {
	return errors.New(s)
}

func Newf(format string, v ...interface{}) error {
	return fmt.Errorf(format, v...)
}

type Option func(*Errors)

func NewMsg(opt ...Option) error {
	err := &Errors{}

	for i := 0; i < len(opt); i++ {
		opt[i](err)
	}

	return err
}

func HTTP(i int) Option {
	return func(err *Errors) {
		err.httpStatus = i
	}
}

func Msg(s string) Option {
	return func(err *Errors) {
		err.Message = s
	}
}

func Msgf(format string, v ...interface{}) Option {
	return func(err *Errors) {
		err.Message = fmt.Sprintf(format, v...)
	}
}

func Type(s string) Option {
	return func(err *Errors) {
		err.Type = s
	}
}

func Code(s string) Option {
	return func(err *Errors) {
		err.Code = s
	}
}

func Dev(s string) Option {
	return func(err *Errors) {
		err.Message = s
	}
}

type Errors struct {
	httpStatus int    `json:"-"`
	Message    string `json:"message,omitempty"`
	Type       string `json:"error,omitempty"`
	Code       string `json:"code,omitempty"`
	DevMsg     string `json:"dev_message,omitempty"`
}

func (err *Errors) Error() string {
	return err.Message
}

func (err *Errors) Http() int {
	return err.httpStatus
}

func IsStatus(status int, err error) bool {
	if errs, ok := err.(*Errors); ok {
		if status == errs.httpStatus {
			return true
		}
	}

	return false
}

func Panic(err error) {
	if err != nil {
		panic(err)
	}
}
