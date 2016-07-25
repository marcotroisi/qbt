package qbt

import (
	"os"
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
	return &CommandsFromJsonFile{}
}

type CommandsFromJsonFile struct {
	filepath string
}

func (c *CommandsFromJsonFile) FindOne(command string) SimpleCommand {

}

func (c *CommandsFromJsonFile) Find(command []string) []SimpleCommand {

}

func (c *CommandsFromJsonFile) openAndParseFile() {
	//open file
	content, err := ioutil.ReadFile(c.filepath)
	if err != nil {
		log.Fatal(err)
	}
	//file content
	m := map[string]map[string]string{}
	er := json.Unmarshal(content, &m)
	if er != nil {
		panic(err)
	}
	//traverse map (this to be done in Find and FindOne
	for college, schools := range m {
		fmt.Println(college, ":")
		for school, id := range schools {
			fmt.Println("\t", school, ":", id)
		}
	}
}