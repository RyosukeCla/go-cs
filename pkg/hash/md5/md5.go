// md5 digest: see rfc 1321 (https://tools.ietf.org/html/rfc1321)

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

// paddings s.t. (message + paddings) modulo 512 == 448
func getPaddingBits(messageLength int) []byte {
	modulo := messageLength % 512
	paddingLength := 512 - modulo
	if paddingLength == 0 {
		paddingLength = 512
	}

	paddings := make([]byte, paddingLength)
	for i := 0; i < paddingLength; i++ {
		paddings[i] = 0
	}
	return paddings
}

func getLowOrderWordsFromLength(length int) []byte {
	littleEndians := make([]byte, 0, 8)
	for i := 0; i < 8; i++ {
		littleEndians = append(littleEndians, byte((length>>(8*i))&0xFF))
	}
	return littleEndians
}

func bytesToWords(bytes []byte) []uint32 {
	length := 8 * len(bytes) / 32
	words := make([]uint32, 0, length)
	for i := 0; i < length; i++ {
		j := i * 4
		words[i] = (uint32(bytes[j+3]) << 24) + (uint32(bytes[j+2]) << 16) + (uint32(bytes[j+1]) << 8) + uint32(bytes[j])
	}
	return words
}

func bitRotateLeft(bits uint32, r int) uint32 {
	return (bits << r) | (bits >> (32 - r))
}

// Hash returns hash. message has 8 * bits length.
func Hash(message []byte) []uint32 {
	// Step 1: append paddings bits
	paddings := getPaddingBits(len(message))
	// Step 2: append length
	lowOrderWords := getLowOrderWordsFromLength(len(message))
	bytes := make([]byte, 0, len(message)+len(paddings)+len(lowOrderWords))
	bytes = append(bytes, message...)
	bytes = append(bytes, paddings...)
	bytes = append(bytes, lowOrderWords...)
	words := bytesToWords(bytes)
	wordsLength := len(words)

	A := uint32(0x67452301)
	B := uint32(0xEFCDAB89)
	C := uint32(0x98BADCFE)
	D := uint32(0x10325476)
	T := createTable()

	for i := 0; i < wordsLength/16; i++ {
		X := words[16*i : 16*i+16]
		tempA := A
		tempB := B
		tempC := C
		tempD := D

		// Round 1
		p1 := func(a *uint32, b, c, d uint32, k, s, i int) {
			*a = b + bitRotateLeft(*a+f1(b, c, d)+X[k]+T[i-1], s)
		}
		p1(&A, B, C, D, 0, 7, 1)
		p1(&D, A, B, C, 1, 12, 2)
		p1(&C, D, A, B, 2, 17, 3)
		p1(&B, C, D, A, 3, 22, 4)
		p1(&A, B, C, D, 4, 7, 5)
		p1(&D, A, B, C, 5, 12, 6)
		p1(&C, D, A, B, 6, 17, 7)
		p1(&B, C, D, A, 7, 22, 8)
		p1(&A, B, C, D, 8, 7, 9)
		p1(&D, A, B, C, 9, 12, 10)
		p1(&C, D, A, B, 10, 17, 11)
		p1(&B, C, D, A, 11, 22, 12)
		p1(&A, B, C, D, 12, 7, 13)
		p1(&D, A, B, C, 13, 12, 14)
		p1(&C, D, A, B, 14, 17, 15)
		p1(&B, C, D, A, 15, 22, 16)

		// Round 2
		p2 := func(a *uint32, b, c, d uint32, k, s, i int) {
			*a = b + bitRotateLeft(*a+f2(b, c, d)+X[k]+T[i-1], s)
		}
		p2(&A, B, C, D, 1, 5, 17)
		p2(&D, A, B, C, 6, 9, 18)
		p2(&C, D, A, B, 11, 14, 19)
		p2(&B, C, D, A, 0, 20, 20)
		p2(&A, B, C, D, 5, 5, 21)
		p2(&D, A, B, C, 10, 9, 22)
		p2(&C, D, A, B, 15, 14, 23)
		p2(&B, C, D, A, 4, 20, 24)
		p2(&A, B, C, D, 9, 5, 25)
		p2(&D, A, B, C, 14, 9, 26)
		p2(&C, D, A, B, 3, 14, 27)
		p2(&B, C, D, A, 8, 20, 28)
		p2(&A, B, C, D, 13, 5, 29)
		p2(&D, A, B, C, 2, 9, 30)
		p2(&C, D, A, B, 7, 14, 31)
		p2(&B, C, D, A, 12, 20, 32)

		// Round 3
		p3 := func(a *uint32, b, c, d uint32, k, s, i int) {
			*a = b + bitRotateLeft(*a+f3(b, c, d)+X[k]+T[i-1], s)
		}
		p3(&A, B, C, D, 5, 4, 33)
		p3(&D, A, B, C, 8, 11, 34)
		p3(&C, D, A, B, 11, 16, 35)
		p3(&B, C, D, A, 14, 23, 36)
		p3(&A, B, C, D, 1, 4, 37)
		p3(&D, A, B, C, 4, 11, 38)
		p3(&C, D, A, B, 7, 16, 39)
		p3(&B, C, D, A, 10, 23, 40)
		p3(&A, B, C, D, 13, 4, 41)
		p3(&D, A, B, C, 0, 11, 42)
		p3(&C, D, A, B, 3, 16, 43)
		p3(&B, C, D, A, 6, 23, 44)
		p3(&A, B, C, D, 9, 4, 45)
		p3(&D, A, B, C, 12, 11, 46)
		p3(&C, D, A, B, 15, 16, 47)
		p3(&B, C, D, A, 2, 23, 48)

		// Round 4
		p4 := func(a *uint32, b, c, d uint32, k, s, i int) {
			*a = b + bitRotateLeft(*a+f4(b, c, d)+X[k]+T[i-1], s)
		}
		p4(&A, B, C, D, 0, 6, 49)
		p4(&D, A, B, C, 7, 10, 50)
		p4(&C, D, A, B, 14, 15, 51)
		p4(&B, C, D, A, 5, 21, 52)
		p4(&A, B, C, D, 12, 6, 53)
		p4(&D, A, B, C, 3, 10, 54)
		p4(&C, D, A, B, 10, 15, 55)
		p4(&B, C, D, A, 1, 21, 56)
		p4(&A, B, C, D, 8, 6, 57)
		p4(&D, A, B, C, 15, 10, 58)
		p4(&C, D, A, B, 6, 15, 59)
		p4(&B, C, D, A, 13, 21, 60)
		p4(&A, B, C, D, 4, 6, 61)
		p4(&D, A, B, C, 11, 10, 62)
		p4(&C, D, A, B, 2, 15, 63)
		p4(&B, C, D, A, 9, 21, 64)

		A = tempA + A
		B = tempB + B
		C = tempC + C
		D = tempD + D
	}

	return []uint32{A, B, C, D}
}
