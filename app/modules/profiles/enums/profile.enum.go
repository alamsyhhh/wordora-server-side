package enums

import "errors"

var ErrInvalidGender = errors.New("invalid gender, must be 'male' or 'female'")


func ValidateGender(gender string) error {
	if gender != "male" && gender != "female" {
		return ErrInvalidGender
	}
	return nil
}
