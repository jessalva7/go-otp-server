package repository

type Authenticate interface {
	Authenticate( phoneNumber string, otp int ) bool
	RemoveIfPresent( phoneNumber string ) bool
	SaveOTP( phoneNumber string, otp int )
}

type Message interface {
	SendSMS( phoneNumber string, message string )
}