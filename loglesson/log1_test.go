package loglesson

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func TestOutput(t *testing.T) {
	logger := logrus.New()
	file, _ := os.OpenFile("applicationlog.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	logger.SetOutput(file)

	logger.Info("hello loging")
	logger.Warn("hello loging")
	logger.Error("hello loging")
}

func TestFormater(t *testing.T) {
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})

	logger.Info("hello loging")
	logger.Warn("hello loging")
	logger.Error("hello loging")
}

func TestField(t *testing.T) {
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})
	logger.WithField("username", "andy").Info("halo dunia")
}

func TestIsFileWrite(t *testing.T) {
	file :="applicationlog.log"
	fileContent, err := ioutil.ReadFile(file)
	assert.NoError(t, err, "error reading file")
	assert.NotEmpty(t, fileContent, "file is empty")
}

func TestFields(t *testing.T) {
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})
	logger.WithFields(logrus.Fields{
		"username":"andy","status":true,
	}).Info("halo dunia")
}