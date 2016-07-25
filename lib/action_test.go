package qbt

import (
	"testing"
)

func TestAction(t *testing.T) {
	action := NewActionFromString("test unit")
	arguments := action.Arguments()

	if arguments[0] != "test" {
		t.Error("expected first argument to be test, got ", arguments[0])
	}

	if arguments[1] != "unit" {
		t.Error("expected second argument to be unit, got ", arguments[1])
	}
}
