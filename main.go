package main

import (
	"fmt"

	"github.com/kreativecoder/guess-number/game"
)

func main() {
	output, err := game.ValidNumber("25")
	fmt.Println(output, err)
	fmt.Println(output == 25)
}
