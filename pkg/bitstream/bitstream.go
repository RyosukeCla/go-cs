package bitstream

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

// ReadAsBytes reads bits as bytes
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

// ByteToBits return bits slice
func ByteToBits(oneByte uint8) []byte {
	bits := make([]byte, 8, 8)
	for i := uint8(7); i > 0; i-- {
		bit := (oneByte >> i) & 1
		bits[7-i] = bit
	}
	bits[7] = oneByte & 1
	return bits
}

// BitsToByte return byte. requires 8 bits.
func BitsToByte(bits []byte) byte {
	var oneByte byte = 0
	for _, bit := range bits {
		oneByte = (oneByte << 1) + bit
	}
	return oneByte
}
