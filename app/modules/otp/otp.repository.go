package otp

import (
	"database/sql"
	"time"

	"github.com/doug-martin/goqu/v9"
)

type OTPRepository struct {
	db *goqu.Database
}

func NewOTPRepository(db *sql.DB) *OTPRepository {
	return &OTPRepository{db: goqu.New("postgres", db)}
}

func (r *OTPRepository) CreateOTP(otp *UserOTP) error {
	_, err := r.db.Insert("user_otps").Rows(otp).Executor().Exec()
	return err
}

func (r *OTPRepository) GetOTPByUserID(userID string) (*UserOTP, error) {
	var otp UserOTP
	found, err := r.db.From("user_otps").Where(goqu.Ex{"user_id": userID}).ScanStruct(&otp)
	if !found {
		return nil, nil
	}
	return &otp, err
}

func (r *OTPRepository) DeleteOTPByUserID(userID string) error {
	_, err := r.db.Delete("user_otps").Where(goqu.Ex{"user_id": userID}).Executor().Exec()
	return err
}

func (r *OTPRepository) DeleteExpiredOTPs() error {
	_, err := r.db.Delete("user_otps").Where(goqu.C("expired_at").Lt(time.Now())).Executor().Exec()
	return err
}