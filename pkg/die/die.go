package die

import (
	"math/rand"
	"time"
)

type Roller interface {
	Roll() int
	Value() int
}

type Die struct {
	sides int
	value int
	isSet bool
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
	d.value = rand.Intn(d.sides) + 1
	d.isSet = false
	return d.value
}

func (d *Die) Value() int {
	return d.value
}

func (d *Die) Set(value int) int {
	d.value = value
	d.isSet = true
	return d.value
}

func (d *Die) IsSet() bool {
	return d.isSet
}
