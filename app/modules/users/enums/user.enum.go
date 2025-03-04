package enums

import "errors"

var ErrInvalidRole   = errors.New("invalid role, must be 'admin' or 'viewer'")


func ValidateRole(role string) error {
	if role != "admin" && role != "viewer" {
		return ErrInvalidRole
	}
	return nil
}
