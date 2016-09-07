package qbt

import (
	"testing"
	"fmt"
)

func TestCommandsFromJsonFile_FindOne(t *testing.T) {
	commands := NewCommandsFromJsonFile("/Users/marcotroisi/gocode/src/github.com/marcotroisi/qbt/file.json")
	result := commands.FindOne("test")

	for _, command := range result {
		cmd := NewSimpleCommandWithExec(command, "")
		output := cmd.Run()
		fmt.Println(output)
	}
}