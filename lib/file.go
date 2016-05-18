package qbt

type FileInterface interface {
	Blocks() []BlockInterface
	Run() (err error)
}

type TestFile struct {
}

func (t *TestFile) Blocks() []BlockInterface {

}
