package basic

import "testing"

func TestFactorial(t *testing.T) {
	if Factorial(5) != 120 {
		t.Fatal("Error")
	}

	if Factorial(4) != 24 {
		t.Fatal("Error")
	}

	if Factorial(0) != 1 {
		t.Fatal("Error")
	}
}
