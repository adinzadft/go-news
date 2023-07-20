package internal

import (
	"gopkg.in/gomail.v2"
)

// SendEmailOTP sends an email containing the OTP code to the user's email address
func SendEmailOTP(email, otpCode string) error {
	mailer := gomail.NewMessage()
	mailer.SetHeader("From", "firdaustirta13@gmail.com")
	mailer.SetHeader("To", email)
	mailer.SetHeader("Subject", "Email Verification")
	mailer.SetBody("text/plain", "Your OTP code is: "+otpCode)

	dialer := gomail.NewDialer("smtp.gmail.com", 587, "firdaustirta13@gmail.com", "@dinz123")

	if err := dialer.DialAndSend(mailer); err != nil {
		return err
	}

	return nil
}
