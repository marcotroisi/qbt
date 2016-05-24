package qbt

type BlockInterface interface {
	Commands() []CommandInterface
}

type TestBlock struct {
}
