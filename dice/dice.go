package main

import (
	"fmt"

	"newspire.com/die"
)

func main() {
	d1 := die.NewDie(6)
	d2 := die.NewDie(6)

	d1.Roll()
	d2.Roll()
	fmt.Println("Dice", d1.Value, d2.Value)
}
