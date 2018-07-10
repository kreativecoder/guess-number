package game

import (
	"errors"
	"fmt"
	"strconv"
)

const max = 10

func ValidNumber(num string) (int, error) {
	if len(num) == 0 {
		return -1, errors.New("Number cannot be empty")
	}

	result, err := strconv.Atoi(num)
	if err != nil {
		return -1, errors.New(err.Error())
	}

	if result > max {
		return -1, fmt.Errorf("Number greater than %v", max)
	}

	return result, nil
}
