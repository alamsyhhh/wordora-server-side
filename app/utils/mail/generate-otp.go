package mail

import (
	"fmt"
	"time"

	"golang.org/x/exp/rand"
)

func GenerateOTP() (string, time.Time) {
	otpCode := fmt.Sprintf("%04d", rand.Intn(10000))
	expirationTime := time.Now().Add(5 * time.Minute)
	return otpCode, expirationTime
}
