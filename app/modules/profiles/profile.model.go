package profiles

type Profile struct {
	ID          string `db:"id"`
	UserID      string `db:"user_id"`
	FullName    string `db:"full_name"`
	Gender      string `db:"gender"`
	PhoneNumber string `db:"phone_number"`
	Address     string `db:"address"`
}
