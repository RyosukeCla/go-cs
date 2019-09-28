package bit_stream

import (
	"testing"
)

func TestA(t *testing.T) {
	var actual = A()
	var expected = 10
	if actual != expected {
		t.Fatal("failed test")
	}
}
