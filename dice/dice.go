package main

import (
	"fmt"

	"newspire.com/die"
)

func main() {
	var mydice [2]die.Die

	for i := range mydice {
		mydice[i] = *die.NewDie(6)
	}

	for i := 5; i > 0; i-- {
		fmt.Print("Dice ", mydice[0].Roll(), mydice[1].Roll())
		if mydice[0].Value == 1 && mydice[1].Value == 1 {
			fmt.Print(" Snake Eyes!")
		} else if mydice[0].Value == mydice[1].Value {
			fmt.Print(" Doubles!")
		}
		fmt.Print("\n")
	}

}
