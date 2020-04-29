package bitvector

import (
	"fmt"

	math "github.com/RyosukeCla/go-cs/pkg/math/basic"
)

type BitVector struct {
	bigBlocks      []uint64
	bigBlockSize   uint64 // lgN ^ 2
	smallBlocks    []uint64
	smallBlockSize uint64   // lgN / 2
	bitsBlocks     []uint64 // n * 64 bits
}

// New returns succinct bit vector
func New(bits []uint64) *BitVector {
	n := uint64(len(bits) * 64)
	lgN := math.Log(2.0, float64(n))
	bigBlockSize := uint64(lgN * lgN)
	smallBlockSize := uint64(lgN / 2.0)

	bigBlocks := make([]uint64, 0, n/bigBlockSize+1)
	smallBlocks := make([]uint64, 0, n/smallBlockSize+1)

	// init
	count := uint64(0)
	bigRank1 := uint64(0)
	smallRank1 := uint64(0)
	for _, bits64 := range bits {
		for j := 63; j >= 0; j-- {
			if count%bigBlockSize == 0 {
				bigBlocks = append(bigBlocks, bigRank1)
				smallRank1 = 0
			}
			if count%smallBlockSize == 0 {
				smallBlocks = append(smallBlocks, smallRank1)
			}

			bit := (bits64 >> j) & 1
			if bit == 1 {
				bigRank1++
				smallRank1++
			}

			count++
		}
	}

	fmt.Println(n, lgN, bigBlockSize, smallBlockSize, n/bigBlockSize, n/smallBlockSize)
	fmt.Println(bits, bigBlocks, smallBlocks)

	return &BitVector{
		bigBlocks:      bigBlocks,
		bigBlockSize:   bigBlockSize,
		smallBlocks:    smallBlocks,
		smallBlockSize: smallBlockSize,
		bitsBlocks:     bits,
	}
}

// Rank returns |{k in [0...x]: B[k] = q}|
func (bv *BitVector) Rank(q uint64, x uint64) uint64 {
	if q == 1 {
		return bv.rank1(x)
	} else {
		return x - bv.rank1(x)
	}
}

// Access returns B[x]
func (bv *BitVector) Access(x uint64) uint64 {
	// find a, b s.t. x = 64 * a + b and a in N, 0 <= b < 64
	// then B[x] = 8BitsBlocks[a][b]
	a := x / 64
	b := x % 64
	bits := bv.bitsBlocks[a]
	return (bits >> (63 - b)) & 1
}

func (bv *BitVector) rank1(x uint64) uint64 {
	bigIndex := x / bv.bigBlockSize
	bigRank1 := bv.bigBlocks[bigIndex]
	smallIndex := x / bv.smallBlockSize
	smallRank1 := bv.smallBlocks[smallIndex]

	rank1 := uint64(0)

	leftX := smallIndex * bv.smallBlockSize / 64
	left := leftX + bv.smallBlockSize/64

	for i := leftX; i < left; i++ {
		bits := bv.bitsBlocks[i]
		for j := 63; j >= 0; j-- {
			bit := (bits >> j) & 1
			if bit == 1 {
				rank1 += 1
			}
		}
	}

	return bigRank1 + smallRank1 + rank1
}

// Select returns min{k in [0...n): Rank(q, k) = x}
func (bv *BitVector) Select(q uint64, x uint64) uint64 {
	if q == 1 {
		return bv.select1(x)
	} else {
		return x - bv.select1(x)
	}
}

func (bv *BitVector) select1(x uint64) uint64 {
	// bigIndex := uint64(0)
	// bigRank1 := uint64(0)
	// for i, rank1 := range bv.bigBlocks {
	// 	if rank1 > x {
	// 		break
	// 	}
	// 	bigRank1 = rank1
	// 	bigIndex = uint64(i) * bv.bigBlockSize
	// }

	// smallIndex := bigIndex / bv.smallBlockSize
	// smallRank1 := uint64(0)
	// for i := smallIndex; i < smallIndex+bv.smallBlockSize; i++ {
	// 	if bigRank1+bv.smallBlocks[i] > x {
	// 		break
	// 	}
	// 	smallRank1 = bv.smallBlocks[i]
	// 	smallIndex = uint64(i) * bv.smallBlockSize
	// }

	// index := smallIndex

	// for i:= smallIndex; i < smallIndex + bv.smallBlockSize {
	// 	if bigRank1 + smallRank1 +
	// }

	// return bigIndex + smallIndex + index
	return 0.0
}
