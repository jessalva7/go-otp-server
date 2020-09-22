package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jessalva/go-otp-server/pkg/authenticating"
)

type authenticatingHandler struct {
	authService authenticating.Service
}

func NewAuthenticatingHandler( service authenticating.Service ) authenticatingHandler {

	return authenticatingHandler{ service }

}

func (authHandler authenticatingHandler) AuthenticateOTP( ctx *gin.Context ) {

	var mobileNumberWithOTP struct{

		MobileNumber string `json:"number"`
		Otp          int   `json:"otp"`

	}
	err := ctx.BindJSON( &mobileNumberWithOTP )
	if err != nil {

		ctx.JSON( 400, "Bad Request: Mobile number should be integer" )
		return

	}

	err = authHandler.authService.Authenticate( mobileNumberWithOTP.MobileNumber, mobileNumberWithOTP.Otp )

	if err != nil {

		ctx.JSON( 401, "Authentication failed")
		return

	}

	ctx.JSON(200, fmt.Sprintf("Successfully authenticated : %s",mobileNumberWithOTP.MobileNumber))


}