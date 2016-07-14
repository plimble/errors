package errors

import (
	"fmt"
	"io"
)

type _errorh struct {
	status int
	msg    string
	*stack
}

func (e _errorh) Error() string { return e.msg }

func (e _errorh) Status() int { return e.status }

func (e _errorh) Format(s fmt.State, verb rune) {
	switch verb {
	case 'v':
		if s.Flag('+') {
			io.WriteString(s, e.msg)
			fmt.Fprintf(s, "%+v", e.StackTrace())
			return
		}
		fallthrough
	case 's':
		io.WriteString(s, e.msg)
	}
}

// New returns an error with the supplied message.
func Newh(status int, message string) error {
	return _errorh{
		status,
		message,
		callers(),
	}
}

// Errorf formats according to a format specifier and returns the string
// as a value that satisfies error.
func Errorhf(status int, format string, args ...interface{}) error {
	return _errorh{
		status,
		fmt.Sprintf(format, args...),
		callers(),
	}
}
