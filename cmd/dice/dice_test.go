package main

import (
	"strings"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/newspire/godice/pkg/die/mocks"
)

func TestDiceMessage(t *testing.T) {

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	d1 := mocks.NewMockRoller(mockCtrl)
	d2 := mocks.NewMockRoller(mockCtrl)

	d1.EXPECT().Value().Return(2).AnyTimes()
	d2.EXPECT().Value().Return(2).AnyTimes()
	result := GetDiceMessage(d1, d2)

	if !strings.Contains(result, "Doubles") {
		t.Errorf("Double Test Failed, got %s", result)
	}

	d1 = mocks.NewMockRoller(mockCtrl)
	d2 = mocks.NewMockRoller(mockCtrl)
	d1.EXPECT().Value().Return(1).AnyTimes()
	d2.EXPECT().Value().Return(1).AnyTimes()
	result = GetDiceMessage(d1, d2)

	if !strings.Contains(result, "Snake Eyes") {
		t.Errorf("Snake Eyes Test Failed, got %s", result)
	}

	d1 = mocks.NewMockRoller(mockCtrl)
	d2 = mocks.NewMockRoller(mockCtrl)
	d1.EXPECT().Value().Return(6).AnyTimes()
	d2.EXPECT().Value().Return(6).AnyTimes()
	result = GetDiceMessage(d1, d2)

	if !strings.Contains(result, "Boxcars") {
		t.Errorf("Boxcars Test Failed, got %s", result)
	}

}
