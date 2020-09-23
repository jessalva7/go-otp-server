package repository

type Authenticate interface {
	Authenticate( phoneNumber string, otp string ) bool
	RemoveIfPresent( phoneNumber string ) bool
	SaveOTP( phoneNumber string, otp string )
}

type Message interface {
	SendSMS( phoneNumber string, message string )
}