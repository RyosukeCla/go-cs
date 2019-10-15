package huffman

import (
	"sort"
)

// Symbol type aiueo
type Symbol struct {
	Symbol     byte
	Occurrence uint64
}

// Code type aiueo
type Code struct {
	Symbol byte
	Code   []byte
}

func buildCode(order int, maxOrder int) []byte {
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
	sort.SliceStable(symbols, func(i, j int) bool { return symbols[i].Occurrence > symbols[j].Occurrence })

	codes := make([]Code, symbolLen)
	for i, symbol := range symbols {
		codes[i] = Code{
			Symbol: symbol.Symbol,
			Code:   buildCode(i, symbolLen-1),
		}
	}

	return codes
}
