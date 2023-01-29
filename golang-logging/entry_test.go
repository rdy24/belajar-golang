package golang_logging

import (
	"testing"

	"github.com/sirupsen/logrus"
)

func TestEntry(t *testing.T) {
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})

	logger.Info("Hello World")
	logger.WithField("username", "bar").Info("Hello World")

	entry := logrus.NewEntry(logger)
	entry.WithField("username", "bar").Info("Hello World")
}
