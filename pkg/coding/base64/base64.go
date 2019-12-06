package base64

import (
	"io"

	"github.com/RyosukeCla/go-cs/pkg/bits"
	"github.com/RyosukeCla/go-cs/pkg/math/basic"
)

var encodingTable = map[byte]byte{
	0: []byte("A")[0], 1: []byte("B")[0], 2: []byte("C")[0], 3: []byte("D")[0], 4: []byte("E")[0], 5: []byte("F")[0], 6: []byte("G")[0], 7: []byte("H")[0],
	8: []byte("I")[0], 9: []byte("J")[0], 10: []byte("K")[0], 11: []byte("L")[0], 12: []byte("M")[0], 13: []byte("N")[0], 14: []byte("O")[0],
	15: []byte("P")[0], 16: []byte("Q")[0], 17: []byte("R")[0], 18: []byte("S")[0], 19: []byte("T")[0], 20: []byte("U")[0], 21: []byte("V")[0],
	22: []byte("W")[0], 23: []byte("X")[0], 24: []byte("Y")[0], 25: []byte("Z")[0], 26: []byte("a")[0], 27: []byte("b")[0], 28: []byte("c")[0],
	29: []byte("d")[0], 30: []byte("e")[0], 31: []byte("f")[0], 32: []byte("g")[0], 33: []byte("h")[0], 34: []byte("i")[0], 35: []byte("j")[0],
	36: []byte("k")[0], 37: []byte("l")[0], 38: []byte("m")[0], 39: []byte("n")[0], 40: []byte("o")[0], 41: []byte("p")[0], 42: []byte("q")[0],
	43: []byte("r")[0], 44: []byte("s")[0], 45: []byte("t")[0], 46: []byte("u")[0], 47: []byte("v")[0], 48: []byte("w")[0], 49: []byte("x")[0],
	50: []byte("y")[0], 51: []byte("z")[0], 52: []byte("0")[0], 53: []byte("1")[0], 54: []byte("2")[0], 55: []byte("3")[0], 56: []byte("4")[0],
	57: []byte("5")[0], 58: []byte("6")[0], 59: []byte("7")[0], 60: []byte("8")[0], 61: []byte("9")[0], 62: []byte("+")[0], 63: []byte("/")[0],
}

var decodingTable = makeReverseTable(encodingTable)

func makeReverseTable(table map[byte]byte) map[byte]byte {
	reverseTable := map[byte]byte{}

	for key, value := range table {
		reverseTable[value] = key
	}

	return reverseTable
}

func encodingLength(pLen int) int {
	ceil := int(basic.Ceil(float64(8*pLen) / float64(6)))
	return ceil + 4 - ceil%4
}

// Encode ...
func Encode(p []byte) []byte {
	length := encodingLength(len(p))
	encoded := make([]byte, length)
	index := 0

	b := bits.NewBitStream([]byte{})
	b.WriteFromBytes(p)

	for {
		sixBits := make([]byte, 6)
		n, err := b.Read(sixBits)

		if err == io.EOF {
			break
		}

		for j := n; j < 6; j++ {
			sixBits[j] = 0
		}

		oneByte := bits.ToByte(sixBits)
		encoded[index] = encodingTable[oneByte]
		index++
	}

	for i := index; i < length; i++ {
		encoded[i] = 61 // []byte("=")[0]
	}

	return encoded
}

// Decode ...
func Decode(p []byte) []byte {
	b := bits.NewBitStream([]byte{})
	for _, ele := range p {
		if ele == 61 {
			break
		}
		sixBits := bits.FromByte(decodingTable[ele])[2:8]
		b.Write(sixBits)
	}

	res := make([]byte, b.Len()/8)
	index := 0
	for {
		eightBits := make([]byte, 8)
		n, err := b.Read(eightBits)
		if n != 8 || err == io.EOF {
			break
		}

		res[index] = bits.ToByte(eightBits)
		index++
	}

	return res
}
