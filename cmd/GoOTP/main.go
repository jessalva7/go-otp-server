package main

import (
	"github.com/dgraph-io/ristretto"
	"github.com/gin-gonic/gin"
	"github.com/jessalva/go-otp-server/pkg/authenticating"
	"github.com/jessalva/go-otp-server/pkg/generating"
	"github.com/jessalva/go-otp-server/pkg/handlers"
	"github.com/jessalva/go-otp-server/pkg/repository"
	"log"
	"os"
	"time"
)

func main(){

	server := gin.Default()

	ristrettoRepo := repository.NewRistrettoRepository( &ristretto.Config{
		NumCounters: 1e7,
		MaxCost:     1e7,
		BufferItems: 20,
	}, time.Minute )
	messageRepo := repository.NewMessageRepository(os.Getenv("TWILIO_URL") + os.Getenv("TWILIO_SID") + "/Messages.json")
	
	generateService := generating.NewService(ristrettoRepo, messageRepo)
	authenticateService := authenticating.NewAuthenticateService(ristrettoRepo)
	genHandler := handlers.NewGeneratingHandler( generateService )
	authHandler := handlers.NewAuthenticatingHandler( authenticateService )

	server.GET("/", genHandler.GenerateOTP )
	server.POST("/authenticate", authHandler.AuthenticateOTP)

	if err := server.Run( ":"+os.Getenv("OTP_SERVER_PORT") ); err != nil {

		log.Fatal( err )

	}

}
