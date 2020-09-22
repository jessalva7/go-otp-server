package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/jessalva/go-otp-server/pkg/generating"
)

type generatingHandler struct {
	genService generating.Service
}

func NewGeneratingHandler( service generating.Service ) generatingHandler {

	return generatingHandler{ service }

}

func (genHandler generatingHandler) GenerateOTP( ctx *gin.Context ) {

	var mobileNumber string

	err := ctx.BindJSON( &mobileNumber )
	if err != nil {

		ctx.JSON( 400, "Bad Request: Mobile number should be integer" )
		return

	}

	ctx.JSON(200, genHandler.genService.GenerateOTP( mobileNumber ) )


}