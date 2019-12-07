package md5

import "github.com/RyosukeCla/go-cs/pkg/math/basic"

func createTable() []uint32 {
	table := make([]uint32, 64)
	for i := 1; i < 64; i++ {
		table[i-1] = uint32(basic.Floor(4294967296.0 * basic.Abs(basic.Sin(float64(i)))))
	}
	return table
}

func f1(x, y, z uint32) uint32 {
	return (x & y) | (^x & z)
}

func f2(x, y, z uint32) uint32 {
	return (x & z) | (y & ^z)
}

func f3(x, y, z uint32) uint32 {
	return x ^ y ^ z
}

func f4(x, y, z uint32) uint32 {
	return y ^ (x | ^z)
}

func Encode() {
	A := uint32(0x67452301)
	B := uint32(0xEFCDAB89)
	C := uint32(0x98BADCFE)
	D := uint32(0x10325476)
}
