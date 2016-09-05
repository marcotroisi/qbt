package qbt

import (
	"log"
	"io/ioutil"
	"encoding/json"
	"fmt"
)

type CommandsInterface interface {
	FindOne(action string) []string
	Find(actions []string) []string
}

// Creates a new CommandsFromJson, taking as an argument the path of the json file
func NewCommandsFromJsonFile(filepath string) *CommandsFromJsonFile {
	return &CommandsFromJsonFile{filepath:filepath}
}

type CommandsFromJsonFile struct {
	filepath string
}

func (c *CommandsFromJsonFile) FindOne(action string) []string {
	parsedFile := c.parsedFile(c.fileContent())
	result := []string{}

	for block, commands := range parsedFile {
		if block == action {
			for name, cmd := range commands {
				fmt.Println("- ", name, ":", cmd)
				result = append(result, cmd)
				// here create array of strings
				// with all the commands part of the "action"
			}
		}
	}

	return result
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