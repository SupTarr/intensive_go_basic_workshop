package password

import "errors"

func ValidatePassword(password string) error {
	number := 0
	capital := 0
	special := 0

	if len(password) < 8 {
		return errors.New("the password must be at least 8 characters")
	}

	for _, v := range password {
		switch {
		case 48 <= v && v <= 57:
			number += 1
		case 58 <= v && v <= 64:
			special += 1
		case 65 <= v && v <= 90:
			capital += 1
		}
	}

	if number < 2 {
		return errors.New("the password must contain at least 2 numbers")
	} else if capital < 1 {
		return errors.New("the password must contain at least one capital letter")
	} else if special < 1 {
		return errors.New("the password must contain at least one special character")
	}

	return nil
}
