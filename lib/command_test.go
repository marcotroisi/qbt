package qbt

import (
	"testing"
)

func TestRun(t *testing.T) {
	cmd := NewSimpleCommand("echo", "ciao", &FkExec{})
	result := cmd.Run()
	if result != "echo ciao" {
		t.Error("expected echo ciao, got ", result)
	}
}
