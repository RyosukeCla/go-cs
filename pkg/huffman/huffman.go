package huffman

import (
	"github.com/RyosukeCla/go-playground/pkg/bitstream"
)

type Symbol struct {
	Symbol []byte
	Occurrence uint64
}

type Code struct {
	Symbol []byte
	Code *bits.BitStream
}

// Build
func Build(symbols: []Symbol) []Code {

}
