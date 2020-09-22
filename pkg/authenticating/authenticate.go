package authenticating

import (
	"errors"
	"github.com/jessalva/go-otp-server/pkg/repository"
)

type Service interface {
	Authenticate( phoneNumber string, otp int ) error
}

type authenticateService struct {
	repo repository.Authenticate
}

func NewAuthenticateService( repo repository.Authenticate) Service {
	return &authenticateService{ repo: repo}
}

func (authService *authenticateService ) Authenticate( phoneNumber string, otp int ) error {

	if authService.repo.Authenticate( phoneNumber, otp ) {

		return nil

	}

	return errors.New("authentication failed | Invalid OTP")

}