package bits

import (
	"bytes"
	"errors"
)

// BitStream struct
type BitStream struct {
	bitBuffer *bytes.Buffer
}

// NewBitStream returns BitStream
func NewBitStream(bits []byte) BitStream {
	return BitStream{
		bitBuffer: bytes.NewBuffer(bits),
	}
}

// Write writes bits []{0, 1}
func (bs *BitStream) Write(bits []byte) (int, error) {
	return bs.bitBuffer.Write(bits)
}

// WriteFromBytes converts p to bits and writes bits.
func (bs *BitStream) WriteFromBytes(p []byte) (int, error) {
	n := 0
	for _, oneByte := range p {
		nb, err := bs.Write(FromByte(oneByte))
		if err != nil {
			return n, err
		}
		n += nb
	}
	return n, nil
}

// WritePadds padds zeros
func (bs *BitStream) WritePadds() (int, error) {
	padding := 8 - (bs.bitBuffer.Len() % 8)
	buffer := make([]byte, padding, padding)
	for i := 0; i < padding; i++ {
		buffer[i] = 0
	}
	return bs.Write(buffer)
}

// Read reads bits
func (bs *BitStream) Read(p []byte) (int, error) {
	return bs.bitBuffer.Read(p)
}

// ReadAsBytes reads 8x bits as x bytes
func (bs *BitStream) ReadAsBytes(p []byte) (int, error) {
	length := len(p)
	buffer := make([]byte, 8, 8)
	n := 0
	for i := 0; i < length; i++ {
		nb, err := bs.Read(buffer)
		n += nb
		if err != nil {
			return n, err
		}
		if nb%8 != 0 {
			return n, errors.New("require 8 times bits")
		}
		p[i] = ToByte(buffer)
	}
	return n, nil
}

// Bits returns a slice of bs.Len() bits
func (bs *BitStream) Bits() []byte {
	return bs.bitBuffer.Bytes()
}

// Len returns the number of bits
func (bs *BitStream) Len() int {
	return bs.bitBuffer.Len()
}

// FromByte convert 1 byte to 8 bits
func FromByte(oneByte byte) []byte {
	bits := make([]byte, 8, 8)
	for i := byte(7); i > 0; i-- {
		bit := (oneByte >> i) & 1
		bits[7-i] = bit
	}
	bits[7] = oneByte & 1
	return bits
}

// ToByte convert 8 bits to 1 byte
func ToByte(bits []byte) byte {
	var oneByte byte = 0
	for _, bit := range bits {
		oneByte = (oneByte << 1) + bit
	}
	return oneByte
}
