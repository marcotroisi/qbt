package qbt

import (
	"testing"
)

func TestAction(t *testing.T) {
	action := NewActionFromString("test unit")
	arguments := action.Arguments()

	// run tests here...
}
