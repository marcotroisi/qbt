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
	for _, command := range result {
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

	phpunitExists := false
	for _, command := range result {
		if command == "vendor/bin/phpunit -c phpunit" {
			phpunitExists = true
		}
	}
	if !phpunitExists {
		t.Error("expected \"vendor/bin/phpunit -c phpunit\"")
	}
}