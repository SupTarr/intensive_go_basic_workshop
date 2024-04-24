package thai_id

import (
	"errors"
	"strconv"
	"strings"
)

func Validate(id string) error {
	splitedId := strings.Split(id, "-")
	newId := strings.Join(splitedId, "")
	if len(newId) != 13 {
		return errors.New("ID is not 13 digits")
	}

	lastDigit := 0
	sum := 0
	for i := 0; i < len(newId); i++ {
		digit, err := strconv.Atoi(string(newId[i]))
		if err != nil {
			return errors.New("cannot parse to integer")
		}

		value := digit * (len(newId) - i)
		if i == len(newId)-1 {
			lastDigit = digit
		} else {
			sum += value
		}
	}

	if (11-(sum%11))%10 != lastDigit {
		return errors.New("not pass the validation")
	}

	return nil
}
