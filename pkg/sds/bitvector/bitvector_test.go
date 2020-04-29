package bitvector

import (
	"testing"
)

func TestAccess(t *testing.T) {
	bv := New([]uint64{
		3038287259199220266, // 00101010 00101010 00101010 00101010 00101010 00101010 00101010 00101010
		578721382704613384,  // 00001000 00001000 00001000 00001000 00001000 00001000 00001000 00001000
	})

	if bv.Access(2) != 1 {
		t.Fatalf("error actual: %d, expected: 1", bv.Access(2))
	}

	if bv.Access(12) != 1 {
		t.Fatalf("error actual: %d, expected: 1", bv.Access(12))
	}

	if bv.Access(13) != 0 {
		t.Fatalf("error actual: %d, expected: 0", bv.Access(13))
	}
}

func TestRank(t *testing.T) {
	bv := New([]uint64{
		3038287259199220266, // 00101010 00101010 00101010 00101010 00101010 00101010 00101010 00101010
		578721382704613384,  // 00001000 00001000 00001000 00001000 00001000 00001000 00001000 00001000
	})

	if bv.Rank(1, 3) != 1 {
		t.Fatalf("error actual: %d, expected: 1", bv.Rank(1, 3))
	}

	if bv.Rank(0, 3) != 2 {
		t.Fatalf("error actual: %d, expected: 2", bv.Rank(0, 3))
	}
}

// func TestSelect(t *testing.T) {
// 	bv := New([]uint64{
// 		3038287259199220266, // 00101010 00101010 00101010 00101010 00101010 00101010 00101010 00101010
// 		578721382704613384,  // 00001000 00001000 00001000 00001000 00001000 00001000 00001000 00001000
// 	})

// 	if bv.Select(1, 3) != 7 {
// 		t.Fatalf("error actual: %d, expected: 7", bv.Select(1, 3))
// 	}

// 	if bv.Select(0, 3) != 4 {
// 		t.Fatalf("error actual: %d, expected: 4", bv.Select(0, 3))
// 	}
// }
