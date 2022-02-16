package main

import (
	"fmt"

	"newspire.org/die"
)

func main() {

	var mydice [2]die.Die

	for i := range mydice {
		mydice[i] = *die.NewDie(6)
	}

	for i := 5; i > 0; i-- {
		mydice[0].Roll()
		mydice[1].Roll()
		result := GetDiceMessage(mydice)
		fmt.Println(result)
	}

}

func GetDiceMessage(mydice [2]die.Die) string {
	message := ""

	if mydice[0].Value() == 1 && mydice[1].Value() == 1 {
		message = "Snake Eyes!"
	} else if mydice[0].Value() == 6 && mydice[1].Value() == 6 {
		message = "Boxcars!"
	} else if mydice[0].Value() == mydice[1].Value() {
		message = "Doubles!"
	}
	result := fmt.Sprintf("Dice %d %d %s", mydice[0].Value(), mydice[1].Value(), message)

	return result
}
