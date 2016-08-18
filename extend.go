package errors

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
)

type _errorh struct {
	ErrStatus  int    `json:"status"`
	ErrMessage string `json:"message"`
	ErrCode    string `json:"code"`
	*stack     `json:"-"`
}

func (e _errorh) Error() string { return e.ErrMessage }

func (e _errorh) Status() int { return e.ErrStatus }

func (e _errorh) Code() string { return e.ErrCode }

func (e _errorh) JSONError() error {
	b, _ := json.Marshal(e)
	return errors.New(string(b))
}

func (e _errorh) Format(s fmt.State, verb rune) {
	switch verb {
	case 'v':
		if s.Flag('+') {
			io.WriteString(s, e.ErrMessage)
			fmt.Fprintf(s, "%+v", e.StackTrace())
			return
		}
		fallthrough
	case 's':
		io.WriteString(s, e.ErrMessage)
	}
}

// New returns an error with the supplied message.
func Newh(status int, code, message string) error {
	return _errorh{
		status,
		message,
		code,
		callers(),
	}
}

// Errorf formats according to a format specifier and returns the string
// as a value that satisfies error.
func Errorhf(status int, code, format string, args ...interface{}) error {
	return _errorh{
		status,
		fmt.Sprintf(format, args...),
		code,
		callers(),
	}
}

func ParseJSON(err string) _errorh {
	e := _errorh{}
	errr := json.Unmarshal([]byte(err), &e)
	if errr != nil {
		e.ErrStatus = 500
		e.ErrCode = "JSON"
		e.ErrMessage = errr.Error()
		e.stack = callers()
	}

	e.stack = callers()
	return e
}

func BadRequest(code, message string) error {
	return &_errorh{
		400,
		message,
		code,
		callers(),
	}
}

func Unauthorized(code, message string) error {
	return &_errorh{
		401,
		message,
		code,
		callers(),
	}
}

func Forbidden(code, message string) error {
	return &_errorh{
		403,
		message,
		code,
		callers(),
	}
}

func NotFound(code, message string) error {
	return &_errorh{
		404,
		message,
		code,
		callers(),
	}
}

func InternalServerError(code, message string) error {
	return &_errorh{
		500,
		message,
		code,
		callers(),
	}
}
