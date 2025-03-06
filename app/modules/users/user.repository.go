package users

import (
	"database/sql"
	"wordora/app/modules/users/model"

	"github.com/doug-martin/goqu/v9"
)

type UserRepository struct {
	db *goqu.Database
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: goqu.New("postgres", db)}
}

func (r *UserRepository) CreateUser(user *model.User) error {
	_, err := r.db.Insert("users").Rows(user).Executor().Exec()
	return err
}

func (r *UserRepository) GetUserByEmail(email string) (*model.User, error) {
	var user model.User
	found, err := r.db.From("users").Where(goqu.Ex{"email": email}).ScanStruct(&user)
	if !found {
		return nil, nil
	}
	return &user, err
}

func (r *UserRepository) UpdateUser(user *model.User) error {
    _, err := r.db.Update("users").
        Set(user).
        Where(goqu.Ex{"id": user.ID}).
        Executor().Exec()
    return err
}

func (r *UserRepository) GetUserByID(id string) (*model.User, error) {
	var user model.User
	found, err := r.db.From("users").Where(goqu.Ex{"id": id}).ScanStruct(&user)
	if !found {
		return nil, nil
	}
	return &user, err
}