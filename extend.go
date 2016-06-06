package errors

import (
	"errors"
	"fmt"
)

type httpContext struct {
	status int
	code   string
}

func (h *httpContext) Context() (int, string) {
	return h.status, h.code
}

func Newh(status int, code, text string) error {
	return struct {
		error
		*stack
		*httpContext
	}{
		errors.New(text),
		callers(),
		&httpContext{
			status,
			code,
		},
	}
}

type LogFunc func(args ...interface{})

func LogError(err error, logFunc LogFunc) {
	stack := ""
	type location interface {
		Location() (string, int)
	}
	type message interface {
		Message() string
	}

	for err != nil {
		if err, ok := err.(location); ok {
			file, line := err.Location()
			stack = fmt.Sprintf("%s:%d: ", file, line)
		}
		switch err := err.(type) {
		case message:
			logFunc(stack, err.Message())
		default:
			logFunc(stack, err.Error())
		}

		cause, ok := err.(causer)
		if !ok {
			break
		}
		err = cause.Cause()
	}
}

func Sprint(err error) string {
	stack := ""
	type location interface {
		Location() (string, int)
	}
	type message interface {
		Message() string
	}

	for err != nil {
		if err, ok := err.(location); ok {
			file, line := err.Location()
			stack += fmt.Sprintf("%s:%d: ", file, line)
		}
		switch err := err.(type) {
		case message:
			stack += fmt.Sprintln(err.Message())
		default:
			stack += fmt.Sprintln(err.Error())
		}

		cause, ok := err.(causer)
		if !ok {
			break
		}
		err = cause.Cause()
	}

	return stack
}
