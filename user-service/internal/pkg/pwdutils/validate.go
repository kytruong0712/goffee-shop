package pwdutils

import (
	"errors"
	"regexp"
	"strings"
)

var (
	// ErrEmptyPassword means password is empty
	ErrEmptyPassword = errors.New("password is empty")
	// ErrInvalidPassword means password is invalid
	ErrInvalidPassword = errors.New("invalid password")
)

// ValidatePassword validates password
func ValidatePassword(pwd string) error {
	pwd = strings.TrimSpace(pwd)
	if pwd == "" {
		return ErrEmptyPassword
	}

	re := regexp.MustCompile(`[a-zA-Z0-9!@#\$%\^&\*]{8,12}`)

	if !re.MatchString(pwd) {
		return ErrInvalidPassword
	}

	return nil
}
