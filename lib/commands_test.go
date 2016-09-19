package qbt

import (
	"testing"
	"path/filepath"
)

func TestCommandsFromJsonFile_FindOne(t *testing.T) {
	absPath, _ := filepath.Abs("../file.json")
	commands := NewCommandsFromJsonFile(absPath)
	result := commands.FindOne("test", "")

	phpspecExists := false
	phpunitExists := false
	for i, command := range result {
		if command == "vendor/bin/phpspec" {
			phpspecExists = true
		}
		if command == "vendor/bin/phpunit -c phpunit" {
			phpunitExists = true
		}
	}

	if !phpspecExists {
		t.Error("expected \"vendor/bin/phpspec\"")
	}
	if !phpunitExists {
		t.Error("expected \"vendor/bin/phpunit -c phpunit\"")
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