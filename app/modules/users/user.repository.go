package users

import (
	"database/sql"
	"log"
	"time"
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
    result, err := r.db.Update("users").
        Set(goqu.Record{
            "is_email_verified": user.IsEmailVerified,
            "updated_at":        time.Now(),
        }).
        Where(goqu.Ex{"id": user.ID}).
        Executor().Exec()

    if err != nil {
        log.Println("Failed to execute update query:", err)
    } else {
        rowsAffected, _ := result.RowsAffected()
        log.Println("Rows affected:", rowsAffected)
    }

    return err
}



func (r *UserRepository) UpdateUserRole(user *model.User) error {
    _, err := r.db.Update("users").
        Set(goqu.Record{"role": user.Role, "updated_at": time.Now()}).
        Where(goqu.Ex{"id": user.ID}).
        Executor().Exec()
    if err != nil {
        log.Println("Failed to execute update query:", err)
    }
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

func (r *UserRepository) GetAllViewerUsers() ([]model.User, error) {
	var users []model.User

	query := r.db.From("users").
		Select("email").
		Where(goqu.C("role").Eq("viewer"))

	err := query.ScanStructs(&users)
	if err != nil {
		log.Println("Error fetching viewer users:", err)
		return nil, err
	}

	return users, nil
}