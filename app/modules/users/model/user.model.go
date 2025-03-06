package model

import "time"

type User struct {
	ID        string    `db:"id"`
	Name      string    `db:"name"`
	Email     string    `db:"email"`
	IsEmailVerified bool  `db:"is_email_verified"`
	Password  string    `db:"password"`
	Role      string    `db:"role"`
	IsAdFree  bool      `db:"is_ad_free"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}