package die

import (
	"math/rand"
	"time"
)

type Die struct {
	sides int
	Value int
}

// init sets initial values for variables used in the function.
func init() {
	rand.Seed(time.Now().UnixNano())
}

func NewDie(sides int) *Die {
	d := new(Die)
	if sides == 0 {
		d.sides = 6
	} else {
		d.sides = sides
	}
	d.Roll()
	return d
}

func (d *Die) Roll() int {
	d.Value = rand.Intn(d.sides) + 1
	return d.Value
}
