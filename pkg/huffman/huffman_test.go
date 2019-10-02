package huffman

import (
	"fmt"
	"testing"
)

func TestBuild(t *testing.T) {
	codes := Build([]Symbol{
		{Symbol: "a", Occurrence: 1},
		{Symbol: "b", Occurrence: 2},
		{Symbol: "c", Occurrence: 3},
		{Symbol: "d", Occurrence: 4},
	})

	fmt.Println(codes)
	if 0 != 0 {
		t.Fatal("aoiep")
	}
}
