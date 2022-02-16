package main

import (
	"strings"
	"testing"

	"github.com/golang/mock/gomock"
	"newspire.org/die"
)

func TestDiceMessages(t *testing.T) {

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	//mockRoll := mock_roll.NewMockRoll(mockCtrl)

	var mydice [2]die.Die

	for i := range mydice {
		mydice[i] = *die.NewDie(6)
	}

	mydice[0].Set(2)
	mydice[1].Set(2)

	result := GetDiceMessage(mydice)

	if !strings.Contains(result, "Doubles") {
		t.Error("Double Test Failed")
	}

	mydice[0].Set(1)
	mydice[1].Set(1)

	result = GetDiceMessage(mydice)

	if !strings.Contains(result, "Snake Eyes") {
		t.Error("Snake Eyes Test Failed")
	}

	mydice[0].Set(6)
	mydice[1].Set(6)

	result = GetDiceMessage(mydice)

	if !strings.Contains(result, "Boxcars") {
		t.Error("Boxcars Test Failed")
	}

}
