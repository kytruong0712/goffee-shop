package phonenoutils

import (
	"errors"
	"regexp"
)

var (
	// ErrInvalidPhoneNumber means phone number is invalid
	ErrInvalidPhoneNumber = errors.New("invalid phone number")
)

// ValidatePhoneNumber validates phone number
func ValidatePhoneNumber(phoneNo string) error {
	var (
		phoneNoLen = 10
		re         = regexp.MustCompile(`^(0|84)(2(0[3-9]|1[0-689]|2[0-25-9]|3[2-9]|4[0-9]|5[124-9]|6[0369]|7[0-7]|8[0-9]|9[012346789])|3[2-9]|5[25689]|7[06-9]|8[0-9]|9[012346789])([0-9]{7})$`)
	)

	if len(phoneNo) < phoneNoLen || !re.MatchString(phoneNo) {
		return ErrInvalidPhoneNumber
	}

	return nil
}
