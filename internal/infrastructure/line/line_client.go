package line

import (
	"errors"

	"github.com/sirupsen/logrus"
)

type LineClient interface {
	Boardcast(message string) error
}

type lineClientImpl struct {
	accessToken string
}

func NewMockedLineClient(accessToken string) LineClient {
	return &lineClientImpl{
		accessToken: accessToken,
	}
}

func (l *lineClientImpl) Boardcast(message string) error {
	if message == "ErrorMsg" {
		return errors.New("failed to send message")
	}

	logrus.Infof("LineClient: Broadcasting message: %s\n", message)

	return nil
}
