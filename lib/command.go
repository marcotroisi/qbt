package qbt

import (
	"log"
	"os/exec"
)

type CommandInterface interface {
	Run() string
}

func NewSimpleCommand(command string, args string, exec ExecInterface) *SimpleCommand {
	return &SimpleCommand{command: command, args: args, exec: exec}
}

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
