package users

import (
	"database/sql"

	"github.com/doug-martin/goqu/v9"
)

type UserRepository struct {
	db *goqu.Database
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: goqu.New("postgres", db)}
}

func (r *UserRepository) CreateUser(user *User) error {
	_, err := r.db.Insert("users").Rows(user).Executor().Exec()
	return err
}

func (r *UserRepository) GetUserByEmail(email string) (*User, error) {
	var user User
	found, err := r.db.From("users").Where(goqu.Ex{"email": email}).ScanStruct(&user)
	if !found {
		return nil, nil
	}
	return &user, err
}

func (r *UserRepository) UpdateUser(user *User) error {
    _, err := r.db.Update("users").
        Set(user).
        Where(goqu.Ex{"id": user.ID}).
        Executor().Exec()
    return err
}
