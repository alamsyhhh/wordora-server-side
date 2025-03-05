package otp

import "time"

type UserOTP struct {
	ID        string    `db:"id"`
	UserID    string    `db:"user_id"`
	OTPCode   string    `db:"otp_code"`
	ExpiredAt time.Time `db:"expired_at"`
	CreatedAt time.Time `db:"created_at"`
}