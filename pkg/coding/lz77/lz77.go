package lz77

import (
	"bytes"
	"io"
)

// Code for lz77
type Code struct {
	Dist  int
	Len   int
	Token []byte
}

// Encoder implements io.ReadWriter interface
type Encoder struct {
	WindowSize      int
	buffer          *bytes.Buffer
	searchBuffer    *bytes.Buffer
	lookAheadBuffer *bytes.Buffer
}

// NewEncoder returns Encoder
func NewEncoder(windowSize int) Encoder {
	return Encoder{
		WindowSize:      windowSize,
		buffer:          bytes.NewBuffer([]byte{}),
		searchBuffer:    bytes.NewBuffer([]byte{}),
		lookAheadBuffer: bytes.NewBuffer([]byte{}),
	}
}

// Write writes buffer
func (en *Encoder) Write(p []byte) (int, error) {
	return en.buffer.Write(p)
}

// Read decodes buffer and reads codes
func (en *Encoder) Read(codes []Code) (int, error) {
	codeLen := len(codes)
	n := 0

	for i := 0; i < codeLen; i++ {
		en.slideLookAheadBuffer()

		if en.lookAheadBuffer.Len() == 0 {
			return n, io.EOF
		}

		n++

		search := en.searchBuffer.Bytes()
		lookAhead := en.lookAheadBuffer.Bytes()
		index := 0
		length := 0

		jStart := en.WindowSize
		if jStart > len(lookAhead) {
			jStart = len(lookAhead)
		}
		for j := jStart; j > 0; j-- {
			partialLookAhead := lookAhead[:j]
			length = j
			index = bytes.LastIndex(search, partialLookAhead)
			if index == -1 {
				continue
			} else {
				break
			}
		}
		if index == -1 {
			token := make([]byte, 1)
			en.lookAheadBuffer.Read(token)
			en.slideSearchBuffer(token)
			codes[i] = Code{
				Dist:  0,
				Len:   0,
				Token: token,
			}
		} else {
			token := make([]byte, length)
			en.lookAheadBuffer.Read(token)
			en.slideSearchBuffer(token)
			codes[i] = Code{
				Dist:  len(search) - index,
				Len:   length,
				Token: []byte{},
			}
		}
	}

	return n, nil
}

func (en *Encoder) slideLookAheadBuffer() {
	length := en.WindowSize - en.lookAheadBuffer.Len()
	bufferForLookAhead := make([]byte, length)
	n, err := en.buffer.Read(bufferForLookAhead)
	if err != nil {
		return
	}
	en.lookAheadBuffer.Write(bufferForLookAhead[:n])
}

func (en *Encoder) slideSearchBuffer(token []byte) {
	en.searchBuffer.Write(token)
	length := en.searchBuffer.Len() - en.WindowSize
	if length < 0 {
		length = 0
	}
	remove := make([]byte, length)
	en.searchBuffer.Read(remove)
}

// Decoder implements io.ReadWriter interface
type Decoder struct {
	buffer *bytes.Buffer
}

// NewDecoder returns Decoder
func NewDecoder() Decoder {
	return Decoder{
		buffer: bytes.NewBuffer([]byte{}),
	}
}

// Write writes and decodes codes
func (de *Decoder) Write(codes []Code) (int, error) {
	n := 0
	for _, code := range codes {
		if code.Len == 0 {
			nb, _ := de.buffer.Write(code.Token)
			n += nb
		} else {
			dictionary := de.buffer.Bytes()
			length := de.buffer.Len()
			from := length - code.Dist
			to := from + code.Len
			nb, _ := de.buffer.Write(dictionary[from:to])
			n += nb
		}
	}

	return n, nil
}

// Read reads decoded codes
func (de *Decoder) Read(p []byte) (int, error) {
	return de.buffer.Read(p)
}
