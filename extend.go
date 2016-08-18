package errors

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

type _errorh struct {
	ErrStatus      string `json:"status"`
	ErrMessage     string `json:"message"`
	ErrMessageCode string `json:"message_code"`
	ErrCode        int    `json:"code"`
	*stack         `json:"-"`
}

func (e _errorh) Error() string { return e.ErrMessageCode + ": " + e.ErrMessage }

func (e _errorh) Status() string { return e.ErrStatus }

func (e _errorh) Code() int { return e.ErrCode }

func (e _errorh) Message() string { return e.ErrMessage }

func (e _errorh) MessageCode() string { return e.ErrMessageCode }

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
func Newh(code int, messageCode, message string) error {
	return _errorh{
		http.StatusText(code),
		message,
		messageCode,
		code,
		callers(),
	}
}

// Errorf formats according to a format specifier and returns the string
// as a value that satisfies error.
func Errorhf(code int, messageCode, format string, args ...interface{}) error {
	return _errorh{
		http.StatusText(code),
		fmt.Sprintf(format, args...),
		messageCode,
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
		e.ErrMessageCode = "JSON"
		e.ErrCode = 500
		e.stack = callers()
	}

	e.stack = callers()
	return e
}

func BadRequest(messageCode, message string) error {
	return &_errorh{
		http.StatusText(400),
		message,
		messageCode,
		400,
		callers(),
	}
}

func Unauthorized(messageCode, message string) error {
	return &_errorh{
		http.StatusText(401),
		message,
		messageCode,
		401,
		callers(),
	}
}

func Forbidden(messageCode, message string) error {
	return &_errorh{
		http.StatusText(403),
		message,
		messageCode,
		403,
		callers(),
	}
}

func NotFound(messageCode, message string) error {
	return &_errorh{
		http.StatusText(404),
		message,
		messageCode,
		404,
		callers(),
	}
}

func InternalServerError(messageCode, message string) error {
	return &_errorh{
		http.StatusText(500),
		message,
		messageCode,
		500,
		callers(),
	}
}
