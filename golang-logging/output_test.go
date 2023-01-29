package golang_logging

import (
	"os"
	"testing"

	"github.com/sirupsen/logrus"
)

func TestOutput(t *testing.T) {
	logger := logrus.New()

	file, _ := os.OpenFile("output.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	logger.SetOutput(file)

	logger.Info("Hello World")
	logger.Warn("Hello World")
	logger.Error("Hello World")
}
