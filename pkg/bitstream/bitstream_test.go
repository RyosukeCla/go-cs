package bitstream

import (
	"fmt"
	"testing"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func TestBitStream_Write(t *testing.T) {
	var bs = NewBitStream()

	var testBits = []uint8{1, 1, 1, 0, 1, 1, 0, 1, 1, 1, 1, 0, 1, 0, 0, 1, 1, 0, 1, 0, 0, 0, 0, 1, 0}

	bs.Write(testBits)

	fmt.Println(bs)
	bs.WritePadds()

	fmt.Printf("%b \n", bs.BitBuffer.Bytes())
	fmt.Printf("%d \n", bs.BitBuffer.Len())

	// file, err := os.Create("./result.txt")
	// check(err)
	// defer file.Close()

	// _, err = file.Write(aiueo)
	// check(err)

	// file.Sync()
	// file.Close()

	if 0 != 0 {
		t.Fatal("failed test")
	}
}
