package qbt

import (
	"testing"
	"path/filepath"
)

func TestCommandsFromJsonFile_FindOne(t *testing.T) {
	absPath, _ := filepath.Abs("../file.json")
	commands := NewCommandsFromJsonFile(absPath)
	result := commands.FindOne("test", "")

	for i, command := range result {
		if i == 0 && command != "vendor/bin/phpspec" {
			t.Error("expected \"vendor/bin/phpspec\", got ", command)
		}
		if i == 1 && command != "vendor/bin/phpunit -c phpunit" {
			t.Error("expected \"vendor/bin/phpunit -c phpunit\", got ", command)
		}
	}
}

func TestCommandsFromJsonFile_FindOneWithBlock(t *testing.T) {
	absPath, _ := filepath.Abs("../file.json")
	commands := NewCommandsFromJsonFile(absPath)
	result := commands.FindOne("test", "unit")

	for i, command := range result {
		if i == 0 && command != "vendor/bin/phpunit -c phpunit" {
			t.Error("expected \"vendor/bin/phpunit -c phpunit\", got ", command)
		}
	}
}