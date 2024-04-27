package addstring

import "testing"

func TestAddEmptyString(t *testing.T) {
	// arrange
	given := ""
	expected := 0

	// act
	actual := AddString(given)

	// assert
	if expected != actual {
		t.Errorf("given %q expected %d but actual %v\n", given, expected, actual)
	}
}

func TestAddOne(t *testing.T) {
	// arrange
	given := "1"
	expected := 1

	// act
	actual := AddString(given)

	// assert
	if expected != actual {
		t.Errorf("given %q expected %d but actual %v\n", given, expected, actual)
	}
}

func TestAddOneAndTwo(t *testing.T) {
	// arrange
	given := "1,2"
	expected := 3

	// act
	actual := AddString(given)

	// assert
	if expected != actual {
		t.Errorf("given %q expected %d but actual %v\n", given, expected, actual)
	}
}

func TestAddOneAndTwoWithNewLines(t *testing.T) {
	// arrange
	given := "1,2\n3"
	expected := 6

	// act
	actual := AddString(given)

	// assert
	if expected != actual {
		t.Errorf("given %q expected %d but actual %v\n", given, expected, actual)
	}
}
