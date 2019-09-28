package bitstream

import (
	"bytes"
	"errors"
	"fmt"
	"io"
)

// BitStream struct
type BitStream struct {
	Padding   int
	BitBuffer *bytes.Buffer
}

// NewBitStream returns BitStream
func NewBitStream() BitStream {
	return BitStream{
		Padding:   0,
		BitBuffer: bytes.NewBuffer([]byte{}),
	}
}

// WriteBit writes bit to buffer
func (bs BitStream) WriteBit(bit uint8) (int, error) {
	if bit != 0 && bit != 1 {
		return 0, errors.New("only 0 or 1")
	}
	return bs.BitBuffer.Write([]byte{bit})
}

// WritePadds padds zeros
func (bs BitStream) WritePadds() (int, error) {
	bs.Padding = 8 - (bs.BitBuffer.Len() % 8)
	var size = 0
	for i := 0; i < bs.Padding; i++ {
		n, err := bs.WriteBit(0)
		size += n
		if err != nil {
			return size, err
		}
	}
	return size, nil
}

// ToByteBuffer returns ByteBuffer
func (bs BitStream) ToByteBuffer() *bytes.Buffer {
	byteBuffer := bytes.NewBuffer([]byte{})
	buffer := make([]uint8, 8)
	for {
		_, err := bs.BitBuffer.Read(buffer)
		if err == io.EOF {
			break
		}
		fmt.Printf("%b \n", buffer)
		var bits uint8 = 0
		for _, bit := range buffer {
			bits = (bits << 1) + bit
		}
		fmt.Printf("%b \n", bits)
		byteBuffer.Write([]byte{bits})
	}
	fmt.Printf("%b \n", byteBuffer)
	return byteBuffer
}
