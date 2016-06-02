package qbt

type FileInterface interface {
	Blocks() []BlockInterface
	Run() (err error)
}

type TestFile struct {
}

func (t *TestFile) Blocks() []BlockInterface {
	return t.structToInterface([]*TestBlock{&TestBlock{}})
}

func (t *TestFile) structToInterface(blockStruct []*TestBlock) []BlockInterface {
	blockInterface := make([]BlockInterface, len(blockStruct))
	for i := range blockStruct {
		blockInterface[i] = blockStruct[i]
	}
	return blockInterface
}
