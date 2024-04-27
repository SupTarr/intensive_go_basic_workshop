package thai_id

import (
	"testing"
)

func Test12InvalidDigits(t *testing.T) {
	// arrange
	given := "190980243647"
	expected := "ID is not 13 digits"

	// act
	actual := Validate(given)

	// assert
	if actual.Error() != expected {
		t.Errorf("given %q expected %s but actual %v\n", expected, actual, actual)
	}
}

func Test13InvalidDigits(t *testing.T) {
	// arrange
	given := "1909802436474"
	expected := "not pass the validation"

	// act
	actual := Validate(given)

	// assert
	if actual.Error() != expected {
		t.Errorf("given %q expected %s but actual %v\n", expected, actual, actual)
	}
}

func Test13ValidDigits(t *testing.T) {
	// arrange
	given := "1909802436470"

	// act
	actual := Validate(given)

	// assert
	if actual != nil {
		t.Errorf("given %q expected %s but actual %v\n", given, "nil", actual)
	}
}

func TestValid13DigitsWithDash(t *testing.T) {
	// arrange
	given := "1-9098-02436-47-0"

	// act
	actual := Validate(given)

	// assert
	if actual != nil {
		t.Errorf("given %q expected %s but actual %v\n", given, "nil", actual)
	}
}
