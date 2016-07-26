package qbt

import (
	"log"
	"io/ioutil"
	"encoding/json"
	"fmt"
)

type CommandsInterface interface {
	FindOne(command string) SimpleCommand
	Find(commands []string) []SimpleCommand
}

// Creates a new CommandsFromJson, taking as an argument the path of the json file
func NewCommandsFromJsonFile(filepath string) *CommandsFromJsonFile {
	return &CommandsFromJsonFile{filepath:filepath}
}

type CommandsFromJsonFile struct {
	filepath string
}

func (c *CommandsFromJsonFile) FindOne(command string) string {
	parsedFile := c.parsedFile(c.fileContent())

	for block, commands := range parsedFile {
		fmt.Println(block, ":")
		for name, command := range commands {
			fmt.Println("\t", name, ":", command)
		}
	}

	return "command"
}

//func (c *CommandsFromJsonFile) Find(command []string) []SimpleCommand {
//
//}

func (c *CommandsFromJsonFile) fileContent() []byte {
	//open file
	content, err := ioutil.ReadFile(c.filepath)
	if err != nil {
		log.Fatal(err)
	}
	return content
}

func (c *CommandsFromJsonFile) parsedFile(fileContent []byte) map[string]map[string]string {
	m := map[string]map[string]string{}
	err:= json.Unmarshal(fileContent, &m)
	if err != nil {
		log.Fatal(err)
	}
	return m
}