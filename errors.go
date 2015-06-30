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

func Http(i int) Option {
	return func(err *Errors) {
		err.HttpStatus = i
	}
}

func Msg(s string) Option {
	return func(err *Errors) {
		err.ErrMessage = s
	}
}

func Msgf(format string, v ...interface{}) Option {
	return func(err *Errors) {
		err.ErrMessage = fmt.Sprintf(format, v...)
	}
}

func Type(s string) Option {
	return func(err *Errors) {
		err.ErrType = s
	}
}

func Code(s string) Option {
	return func(err *Errors) {
		err.ErrCode = s
	}
}

func Dev(s string) Option {
	return func(err *Errors) {
		err.ErrMessage = s
	}
}

type Errors struct {
	HttpStatus int    `json:"-"`
	ErrMessage string `json:"message,omitempty"`
	ErrType    string `json:"error,omitempty"`
	ErrCode    string `json:"code,omitempty"`
	ErrDevMsg  string `json:"dev_message,omitempty"`
}

func (err *Errors) Error() string {
	return err.ErrMessage
}

func (err *Errors) Http(i int) *Errors {
	err.HttpStatus = i

	return err
}

func (err *Errors) Type(s string) *Errors {
	err.ErrType = s

	return err
}

func (err *Errors) Code(s string) *Errors {
	err.ErrCode = s

	return err
}

func (err *Errors) Dev(s string) *Errors {
	err.ErrDevMsg = s

	return err
}

func IsStatus(status int, err error) bool {
	if errs, ok := err.(*Errors); ok {
		if status == errs.HttpStatus {
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
