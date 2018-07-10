package game

import (
	"errors"
	"testing"
)

func TestValidNumber(t *testing.T) {

	var cases = []struct {
		param         string
		expected      int
		expectedError error
	}{
		{"", -1, errors.New("Number cannot be empty")},
		{"9", 9, nil},
		{"10010", -1, errors.New("Number greater than 10")},
		{"xyz", -1, errors.New("Not a number")},
	}

	for _, test := range cases {
		actual, err := ValidNumber(test.param)

		if err != nil {

			if err.Error() != test.expectedError.Error() {
				t.Errorf("Expected: %v, Got: %v", test.expectedError, err)
			}

			if actual != test.expected {
				t.Errorf("Expected ValidNumber(%q) to be %v, got %v", test.param, test.expected, actual)
			}
		} else {
			if actual != test.expected {
				t.Errorf("Expected ValidNumber(%q) to be %v, got %v", test.param, test.expected, actual)
			}
		}

	}
}

func TestNumberRange(t *testing.T) {

	var cases = []struct {
		value    int
		guess    int
		expected string
	}{
		{9, 8, "Too Low"},
		{5, 8, "Too high"},
		{3, 3, "Correct"},
	}

	for _, test := range cases {
		actual := NumberRange(test.value, test.guess)

		if actual != test.expected {
			t.Errorf("Expected NumberRange(%q, %q) to be %v, got %v", test.value, test.guess, test.expected, actual)
		}

	}
}
