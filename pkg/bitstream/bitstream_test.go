package bitstream

import (
	"fmt"
	"os"
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

	for _, bit := range testBits {
		bs.WriteBit(bit)
	}
	fmt.Println(bs.Padding)
	bs.WritePadds()

	// for i := 0; i < 20; i++ {
	// 	bs.WriteBit(uint8(rand.Intn(2)))
	// }

	fmt.Printf("%b \n", bs.BitBuffer.Bytes())
	fmt.Printf("%d \n", bs.BitBuffer.Len())

	aiueo := bs.ToByteBuffer()

	fmt.Printf("%X", aiueo.Bytes())

	file, err := os.Create("./result.txt")
	check(err)
	defer file.Close()

	_, err = file.Write(aiueo.Bytes())
	check(err)

	file.Sync()
	file.Close()

	if 0 != 0 {
		t.Fatal("failed test")
	}
}
