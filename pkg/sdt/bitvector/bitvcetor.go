package bitvector

type BitVector struct {
	denseBlock   []byte // n * 8 bit
	sparseVector []int
}

// New returns succinct bit vector
func New(denseBlock []byte, sparseVector []int) *BitVector {
	return &BitVector{
		denseBlock,
		sparseVector,
	}
}

// Access returns B[x]
func (bv *BitVector) Access(x uint64) byte {
	// find a, b s.t. x = 8 * a + b and a, b in N
	// then B[x] = 8BitsBlocks[a][b]
	a := x / 8
	b := x % 8
	bits := bv.denseBlock[a]

	return (bits >> (7 - b)) & 1
}

// Rank returns |{k in [0...x]: B[k] = q}|
func (bv *BitVector) Rank(q byte, x uint64) uint64 {
	count := uint64(0)
	cursor := uint64(1)
	n := len(bv.denseBlock)
	for i := 0; i < n; i++ {
		bits := bv.denseBlock[i]

		for j := 7; j >= 0; j-- {
			if cursor > x {
				return count
			}
			bit := (bits >> j) & 1

			if q == bit {
				count++
			}
			cursor++
		}
	}

	return count
}

// Select returns min{k in [0...n): Rank(q, k) = x}
func (bv *BitVector) Select(q byte, x uint64) uint64 {
	count := uint64(0)
	cursor := uint64(0)
	n := len(bv.denseBlock)
	for i := 0; i < n; i++ {
		bits := bv.denseBlock[i]

		for j := 7; j >= 0; j-- {
			if count == x {
				return cursor
			}
			bit := (bits >> j) & 1

			if q == bit {
				count++
			}
			cursor++
		}
	}

	return cursor
}
