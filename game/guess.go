package game

import (
	"errors"
	"fmt"
	"strconv"
)

const max = 10

// ValidNumber : checks if string number is valid i.e not empty, not greater than max and its actually a number
func ValidNumber(num string) (int, error) {
	if len(num) == 0 {
		return -1, errors.New("Number cannot be empty")
	}

	result, err := strconv.Atoi(num)
	if err != nil {
		return -1, errors.New("Not a number")
	}

	if result > max {
		return -1, fmt.Errorf("Number greater than %v", max)
	}

	return result, nil
}

//NumberRange : checks range of value against guess
func NumberRange(value, guess int) string {
	diff := guess - value
	if diff < 0 {
		return "Too Low"
	} else if diff > 0 {
		return "Too high"
	} else {
		return "Correct"
	}
}
