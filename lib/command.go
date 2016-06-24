package qbt

import (
	"log"
	"os/exec"
)

type CommandInterface interface {
	Run() string
}

// Constructor for SimpleCommand that accepts any exec library that is
// compatible with ExecInterface
func NewSimpleCommand(command string, args string, exec ExecInterface) *SimpleCommand {
	return &SimpleCommand{command: command, args: args, exec: exec}
}

// Constructor for SimpleCommand that uses Exec
func NewSimpleCommandWithExec(command string, args string) *SimpleCommand {
	osExec := &Exec{}
	return &SimpleCommand{command: command, args: args, exec: osExec}
}

// SimpleCommand takes the command as it is and executes it
type SimpleCommand struct {
	command string
	args    string
	exec    ExecInterface
}

// Runs the command and returns the output
func (t *SimpleCommand) Run() string {
	return t.command + " " + t.args
}

type ExecInterface interface {
	Run(command string, args string) (output []byte, err error)
}

type Exec struct {
}

func (e *Exec) Run(command string, args string) (output []byte, err error) {
	output, err = exec.Command("date").Output()
	if err != nil {
		log.Fatal(err)
	}
	return output, err
}
