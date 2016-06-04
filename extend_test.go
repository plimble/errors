package errors

import (
	"testing"

	"github.com/Sirupsen/logrus"
)

func TestLog(t *testing.T) {
	LogError(New("error!!"), logrus.Info)
}
