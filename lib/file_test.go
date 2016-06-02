package qbt

import (
	// "reflect"
	"testing"
)

func TestBlocks(t *testing.T) {
	testFile := TestFile{}
	blocks := testFile.Blocks()
	// expected := []*TestBlock{&TestBlock{}}
	if blocks == nil {
		t.Error("expected slice, got ", blocks)
	}
	// if reflect.DeepEqual(blocks, expected) == false {
	// 	t.Error("expected ", expected, ", got ", blocks)
	// }
}
