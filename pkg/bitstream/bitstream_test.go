package bitstream

import (
	"bytes"
	"fmt"
	"testing"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func TestBitStream_Write(t *testing.T) {
	var bits = []uint8{1, 1, 1, 0, 1, 1, 0, 1, 1, 1, 1, 0, 1, 0, 0, 1, 1, 0, 1, 0, 0, 0, 0, 1, 0}
	fmt.Println("bits \t\t", bits)
	var bs = NewBitStream()
	bs.Write(bits)
	bs.WritePadds()

	fmt.Println("padded bits \t", bs.BitBuffer.Bytes())

	byteBuffer := bytes.NewBuffer([]byte{})
	buffer := make([]byte, 1)
	for {
		_, err := bs.ReadAsBytes(buffer)
		if err != nil {
			break
		}
		byteBuffer.Write(buffer)
	}

	fmt.Println("bytes \t\t", byteBuffer.Bytes())

	var bs2 = NewBitStream()
	buffer = make([]byte, 1, 1)
	for {
		_, err := byteBuffer.Read(buffer)
		if err != nil {
			break
		}
		bs2.WriteBytesAsBits(buffer)
	}

	fmt.Println("decoded \t", bs2.BitBuffer.Bytes())

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
