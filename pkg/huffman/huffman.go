package huffman

import (
	"sort"

	"github.com/RyosukeCla/go-playground/pkg/bits"
)

// Symbol type aiueo
type Symbol struct {
	Symbol     string
	Occurrence uint64
}

// Code type aiueo
type Code struct {
	Symbol string
	Code   []byte
}

// BuildCode returns bits
func BuildCode(order int, maxOrder int) []byte {
	if order == 0 {
		return []byte{0}
	} else if order == maxOrder {
		len := order
		codes := make([]byte, len)
		for i := 0; i < len; i++ {
			codes[i] = 1
		}
		return codes
	} else {
		len := order + 1
		codes := make([]byte, len)
		for i := 0; i < order; i++ {
			codes[i] = 1
		}
		codes[len-1] = 0
		return codes
	}
}

// Build builds symbols
func Build(symbols []Symbol) []Code {
	symbolLen := len(symbols)
	sortedSymbols := make([]Symbol, symbolLen)
	copy(sortedSymbols, symbols)
	sort.SliceStable(sortedSymbols, func(i, j int) bool { return symbols[i].Occurrence > symbols[j].Occurrence })

	codes := make([]Code, symbolLen)
	for i, symbol := range sortedSymbols {
		bitStream := bits.NewBitStream()
		bitStream.Write([]byte{0})
		codes[i] = Code{
			Symbol: symbol.Symbol,
			Code:   BuildCode(i, symbolLen-1),
		}
	}

	return codes
}
