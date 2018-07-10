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
