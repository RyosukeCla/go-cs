package huffman

import (
	"github.com/RyosukeCla/go-playground/pkg/bits"
)

type Symbol struct {
	Symbol     []byte
	Occurrence uint64
}

type Code struct {
	Symbol []byte
	Code   *bits.BitStream
}

// Build builds symbols
func Build(symbols []Symbol) []Code {
}
