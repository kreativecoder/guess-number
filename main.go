package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/kreativecoder/guess-number/game"
)

const maxGuess = 3

func main() {
	var num string
	fmt.Println("=============GUESS NAME=============")
	fmt.Println("You can guess a number between 1-10")
	scanner := bufio.NewScanner(os.Stdin)
	noOfGuesses := 1
	for num != "q" {
		rand.Seed(time.Now().UnixNano())
		imaginary := rand.Intn(10)

		for noOfGuesses <= maxGuess {
			response := playGame(scanner, &num, &noOfGuesses, imaginary)
			if response == "Correct" {
				noOfGuesses = 1
				fmt.Println("===========================================")
				break
			}

			noOfGuesses++
			if noOfGuesses > maxGuess {
				noOfGuesses = 1
				fmt.Println("Out of guesses")
				fmt.Println("===========================================")
				break
			}
		}
	}
}

func getUserInput(scanner *bufio.Scanner, num *string, noOfGuesses *int) {
	noString := ""
	switch *noOfGuesses {
	case 1:
		noString = "First"
	case 2:
		noString = "Second"
	case 3:
		noString = "Third"
	}

	fmt.Printf("%s Guess: ", noString)
	scanner.Scan()
	*num = scanner.Text()
}

func validateUserInput(numptr string, imaginary int) (string, error) {
	numVal, err := game.ValidNumber(numptr)
	if err != nil {
		return "", err
	}

	result := game.NumberRange(imaginary, numVal)
	return result, nil
}

func playGame(scanner *bufio.Scanner, num *string, noOfGuesses *int, imaginary int) string {
	getUserInput(scanner, num, noOfGuesses)

	result, err := validateUserInput(*num, imaginary)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(result)
	return result
}
