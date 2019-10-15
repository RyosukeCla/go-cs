package runlength

import (
	"bytes"
	"io"
)

// Code is code
type Code struct {
	Symbol byte
	Count  int
}

// Encoder encodes string with runlength method
type Encoder struct {
	leadBuffer *byte
	buffer     *bytes.Buffer
}

// NewEncoder returns Encoder
func NewEncoder() Encoder {
	return Encoder{
		leadBuffer: nil,
		buffer:     bytes.NewBuffer([]byte{}),
	}
}

func (en *Encoder) Write(p []byte) (int, error) {
	return en.buffer.Write(p)
}

func (en *Encoder) Read(code []Code) (int, error) {
	length := len(code)
	n := 0

	symbols := make([]byte, 1, 1)

	if en.leadBuffer == nil {
		_, err := en.buffer.Read(symbols)
		if err == io.EOF {
			return n, io.EOF
		}
		en.leadBuffer = &symbols[0]
	}

	for i := 0; i < length; i++ {
		count := 1

		for {
			_, err := en.buffer.Read(symbols)
			if err == io.EOF {
				code[i] = Code{
					Symbol: *en.leadBuffer,
					Count:  count,
				}
				n++
				return n, io.EOF
			}

			symbol := symbols[0]
			if symbol == *en.leadBuffer {
				count++
				en.leadBuffer = &symbol
			} else {
				code[i] = Code{
					Symbol: *en.leadBuffer,
					Count:  count,
				}
				en.leadBuffer = &symbol
				n++
				break
			}
		}
	}

	return n, nil
}

// Decoder is decoder
type Decoder struct {
	buffer *bytes.Buffer
}

// NewDecoder returns Decoder
func NewDecoder() Decoder {
	return Decoder{
		buffer: bytes.NewBuffer([]byte{}),
	}
}

func (de *Decoder) Write(p []Code) (int, error) {
	n := 0
	for _, code := range p {
		str := make([]byte, code.Count)
		for i := 0; i < code.Count; i++ {
			str[i] = code.Symbol
		}
		de.buffer.Write(str)
		n++
	}

	return n, nil
}

func (de *Decoder) Read(p []byte) (int, error) {
	return de.buffer.Read(p)
}
