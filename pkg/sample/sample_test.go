package sample

import (
	"testing"
)

func TestDouble(t *testing.T) {
	var actual = MultiplyDouble(10)
	var expected = 20
	if actual != expected {
		t.Fatal("failed test")
	}
}
