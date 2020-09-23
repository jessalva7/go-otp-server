package generating

import (
	"fmt"
	"github.com/jessalva/go-otp-server/pkg/repository"
	"math/rand"
	"os"
	"time"
)

type Service interface{
	GenerateOTP( string ) string
}

type service struct {
	authRepo repository.Authenticate
	msgRepo repository.Message
}

func (s *service) GenerateOTP( phoneNumber string ) string {

	rand.Seed( time.Now().Unix() )
	generatedOTPValue := rand.Intn(1e6)
	generatedOTP := fmt.Sprintf("%06d",generatedOTPValue)

	s.authRepo.SaveOTP( phoneNumber, generatedOTP )

	if os.Getenv("TWILIO_SID") == "" && os.Getenv("TWILIO_AUTH_TOKEN") == ""{

		return "Sending SMS failed"

	}

	s.msgRepo.SendSMS(phoneNumber, generatedOTP)

	return fmt.Sprintf( "generated OTP")

}

func NewService( authRepo repository.Authenticate, msgRepo repository.Message) Service{
	return &service{ authRepo: authRepo, msgRepo: msgRepo}
}