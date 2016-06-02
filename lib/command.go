package qbt

// import (
// 	"os/exec"
// )

type CommandInterface interface {
	Run()
}

func NewSimpleCommand(command string, args string) *SimpleCommand {
	return &SimpleCommand{command: command, args: args}
}

// SimpleCommand takes the command as it is and executes it
type SimpleCommand struct {
	command string
	args    string
}

func (t *SimpleCommand) Run() string {
	return t.command + " " + t.args
}

// type ExecInterface interface {
// 	Run(command string, args string)
// }

// type Exec struct {
// }

// func (e *Exec) Run(command string, args string) {
// 	out, err := exec.Command("date").Output()
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	fmt.Printf("The date is %s\n", out)
// }
