package errors

import (
	"fmt"
	"testing"
)

var helloErr = Newh(400, "hello")

func TestLog(t *testing.T) {
	type HttpStatus interface {
		Status() int
	}
	x := Wrap(helloErr, "dddd")
	y := Wrap(x, "wow")

	if st, ok := Cause(y).(HttpStatus); ok {
		fmt.Println("@@@@", st.Status())
	}

	fmt.Printf("%+v", y)
}
