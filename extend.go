package errors

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

type _errorh struct {
	ErrStatus  string `json:"status"`
	ErrMessage string `json:"message"`
	ErrCode    int    `json:"code"`
	*stack     `json:"-"`
}

func (e _errorh) Error() string { return e.ErrMessage }

func (e _errorh) Status() string { return e.ErrStatus }

func (e _errorh) Code() int { return e.ErrCode }

func (e _errorh) Message() string { return e.ErrMessage }

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
func Newh(code int, message string) error {
	return _errorh{
		http.StatusText(code),
		message,
		code,
		callers(),
	}
}

// Errorf formats according to a format specifier and returns the string
// as a value that satisfies error.
func Errorhf(code int, format string, args ...interface{}) error {
	return _errorh{
		http.StatusText(code),
		fmt.Sprintf(format, args...),
		code,
		callers(),
	}
}

func ParseJSON(err string) _errorh {
	e := _errorh{}
	errr := json.Unmarshal([]byte(err), &e)
	if errr != nil {
		e.ErrStatus = http.StatusText(500)
		e.ErrMessage = errr.Error()
		e.ErrCode = 500
		e.stack = callers()
	}

	e.stack = callers()
	return e
}

func BadRequest(message string) error {
	return &_errorh{
		http.StatusText(400),
		message,
		400,
		callers(),
	}
}

func Unauthorized(message string) error {
	return &_errorh{
		http.StatusText(401),
		message,
		401,
		callers(),
	}
}

func Forbidden(message string) error {
	return &_errorh{
		http.StatusText(403),
		message,
		403,
		callers(),
	}
}

func NotFound(message string) error {
	return &_errorh{
		http.StatusText(404),
		message,
		404,
		callers(),
	}
}

func InternalServerError(message string) error {
	return &_errorh{
		http.StatusText(500),
		message,
		500,
		callers(),
	}
}
