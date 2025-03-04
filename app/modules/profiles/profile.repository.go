package profiles

import (
	"database/sql"

	"github.com/doug-martin/goqu/v9"
)

type ProfileRepository struct {
	db *goqu.Database
}

func NewProfileRepository(db *sql.DB) *ProfileRepository {
	return &ProfileRepository{db: goqu.New("postgres", db)}
}

func (r *ProfileRepository) CreateProfile(profile *Profile) error {
	_, err := r.db.Insert("profiles").Rows(profile).Executor().Exec()
	return err
}
