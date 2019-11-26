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

func TestAbs(t *testing.T) {
	if Abs(5) != 5 {
		t.Fatal("Error")
	}

	if Abs(-5) != 5 {
		t.Fatal("Error")
	}
}

func TestSquare(t *testing.T) {
	if Square(5.0) != 25.0 {
		t.Fatal("Error")
	}
}
