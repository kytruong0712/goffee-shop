package pwdutils

import (
	"errors"
	"regexp"
)

var (
	// ErrInvalidPassword means password is invalid
	ErrInvalidPassword = errors.New("invalid password")
)

// ValidatePassword validates password
func ValidatePassword(pwd string) error {
	re := regexp.MustCompile(`[a-zA-Z0-9!@#\$%\^&\*]{8,12}`)

	if !re.MatchString(pwd) {
		return ErrInvalidPassword
	}

	return nil
}
