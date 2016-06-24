package qbt

type BlockInterface interface {
	Commands() []CommandInterface
}

type TestBlock struct {
}

func (t *TestBlock) Commands() []TestCommand {
	return []TestCommand{TestCommand{}}
}
