package qbt

type FileInterface interface {
	Blocks() []BlockInterface
	Run() (err error)
}

type TestFile struct {
}

func (t *TestFile) Blocks() []BlockInterface {
	// @TODO: need to do conversion from slice
	// of TestBlock to slice of BlockInterface
	return []TestBlock{TestBlock{}}
}
