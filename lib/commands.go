package qbt

import (
	"io/ioutil"
	"encoding/json"
	"log"
)

type CommandsInterface interface {
	FindOne(action string, block string) []string
	Find(actions []string) []string
}

// Creates a new CommandsFromJson, taking as an argument the path of the json file
func NewCommandsFromJsonFile(filepath string) *CommandsFromJsonFile {
	return &CommandsFromJsonFile{filepath:filepath}
}

type CommandsFromJsonFile struct {
	filepath string
}

func (c *CommandsFromJsonFile) FindOne(action string, block string) []string {
	fileContent, err := c.fileContent()
	if (err != nil) {
		log.Fatal(err)
	}
	parsedFile, err := c.parsedFile(fileContent)
	if (err != nil) {
		log.Fatal(err)
	}
	result := []string{}

	for actn, commands := range parsedFile {
		if actn == action {
			for name, cmd := range commands {
				if block != "" && block != name {
					continue
				}
				//fmt.Println("- ", name, ":", cmd)
				// array of strings with all the commands
				// part of the "action"
				result = append(result, cmd)
			}
		}
	}

	return result
}

func (c *CommandsFromJsonFile) Find(actions []string) []string {
	result := []string{}
	for _, action := range actions {
		result = append(result, c.FindOne(action, "")...)
	}

	return result
}

func (c *CommandsFromJsonFile) fileContent() ([]byte, error) {
	//open file
	content, err := ioutil.ReadFile(c.filepath)
	if err != nil {
		return nil, err
	}
	return content, nil
}

func (c *CommandsFromJsonFile) parsedFile(fileContent []byte) (map[string]map[string]string, error) {
	m := map[string]map[string]string{}
	err:= json.Unmarshal(fileContent, &m)
	if err != nil {
		return nil, err
	}
	return m, nil
}