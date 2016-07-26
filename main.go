package main

import (
	"fmt"
	"github.com/marcotroisi/qbt/lib"
)

func main() {
	fmt.Println("This is QBT - Quasi Build & Test")

	commands := qbt.NewCommandsFromJsonFile("/Users/marcotroisi/gocode/src/github.com/marcotroisi/qbt/file.json")
	commands.FindOne("test")
}
