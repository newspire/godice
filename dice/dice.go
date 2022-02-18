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
		result := GetDiceMessage(&mydice[0], &mydice[1])
		fmt.Println(result)
	}

}

func GetDiceMessage(d1 die.Roller, d2 die.Roller) string {
	message := ""

	if d1.Value() == 1 && d2.Value() == 1 {
		message = "Snake Eyes!"
	} else if d1.Value() == 6 && d2.Value() == 6 {
		message = "Boxcars!"
	} else if d1.Value() == d2.Value() {
		message = "Doubles!"
	}
	result := fmt.Sprintf("Dice %d %d %s", d1.Value(), d2.Value(), message)

	return result
}
