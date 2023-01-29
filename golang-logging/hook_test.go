package golang_logging

import (
	"fmt"
	"testing"

	"github.com/sirupsen/logrus"
)

type SampleHook struct{}

func (sampleHook *SampleHook) Levels() []logrus.Level {
	return []logrus.Level{logrus.ErrorLevel, logrus.WarnLevel}
}

func (sampleHook *SampleHook) Fire(entry *logrus.Entry) error {
	fmt.Println("SampleHook: ", entry.Message, entry.Level)
	return nil
}

func TestHook(t *testing.T) {
	logger := logrus.New()
	logger.AddHook(&SampleHook{})

	logger.Info("Hello World")
	logger.Warn("Hello World")
}
