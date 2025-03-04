package users

import "time"

type User struct {
	ID        string    `db:"id"`
	Name      string    `db:"name"`
	Email     string    `db:"email"`
	Password  string    `db:"password"`
	Role      string    `db:"role"`
	IsAdFree  bool      `db:"is_ad_free"`
	CreatedAt time.Time `db:"created"`
	UpdatedAt time.Time `db:"updated"`
}

type Profile struct {
	ID          string `db:"id"`
	UserID      string `db:"user_id"`
	FullName    string `db:"full_name"`
	Gender      string `db:"gender"`
	PhoneNumber string `db:"phone_number"`
	Address     string `db:"address"`
}