package bits

import (
	"bytes"
)

// BitStream struct
type BitStream struct {
	BitBuffer *bytes.Buffer
}

// NewBitStream returns BitStream
func NewBitStream() BitStream {
	return BitStream{
		BitBuffer: bytes.NewBuffer([]byte{}),
	}
}

// Write writes bits []{0, 1}
func (bs *BitStream) Write(bits []byte) (int, error) {
	return bs.BitBuffer.Write(bits)
}

// WriteBytesAsBits converts p to bits and writes bits.
func (bs *BitStream) WriteBytesAsBits(p []byte) (int, error) {
	n := 0
	for _, oneByte := range p {
		nb, err := bs.Write(ByteToBits(oneByte))
		if err != nil {
			return n, err
		}
		n += nb
	}
	return n, nil
}

// WritePadds padds zeros
func (bs *BitStream) WritePadds() (int, error) {
	padding := 8 - (bs.BitBuffer.Len() % 8)
	buffer := make([]byte, padding, padding)
	for i := 0; i < padding; i++ {
		buffer[i] = 0
	}
	return bs.Write(buffer)
}

// ReadAsBytes reads 8x bits as x bytes
func (bs *BitStream) ReadAsBytes(p []byte) (int, error) {
	length := len(p)
	buffer := make([]byte, 8, 8)
	n := 0
	for i := 0; i < length; i++ {
		nb, err := bs.BitBuffer.Read(buffer)
		if err != nil {
			return n, err
		}
		n += nb
		p[i] = BitsToByte(buffer)
	}
	return n, nil
}

// Read reads bits
func (bs *BitStream) Read(p []byte) (int, error) {
	return bs.BitBuffer.Read(p)
}

// Bits returns a slice of bs.Len() bits
func (bs *BitStream) Bits() []byte {
	return bs.BitBuffer.Bytes()
}

// Len returns the number of bits
func (bs *BitStream) Len() int {
	return bs.BitBuffer.Len()
}

// ByteToBits convert 1 byte to 8 bits
func ByteToBits(oneByte byte) []uint8 {
	bits := make([]byte, 8, 8)
	for i := byte(7); i > 0; i-- {
		bit := (oneByte >> i) & 1
		bits[7-i] = bit
	}
	bits[7] = oneByte & 1
	return bits
}

// BitsToByte convert 8 bits to 1 byte
func BitsToByte(bits []uint8) byte {
	var oneByte byte = 0
	for _, bit := range bits {
		oneByte = (oneByte << 1) + bit
	}
	return oneByte
}
