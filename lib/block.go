package qbt

type BlockInterface interface {
	Commands() []TestCommand
}

type TestBlock struct {
}

func (t *TestBlock) Commands() []TestCommand {
	return []TestCommand{TestCommand{}}
}
