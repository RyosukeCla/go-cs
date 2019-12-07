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

func TestCeil(t *testing.T) {
	if Ceil(5.5) != 6.0 {
		t.Fatal("Error")
	}
}

func TestSin(t *testing.T) {
	if Abs(Sin(2.0)-0.90929742682) > 0.0000001 {
		t.Fatal("Error", Sin(2.0))
	}
}

func TestCos(t *testing.T) {
	if Abs(Cos(1.0)-0.54030230586) > 0.0000001 {
		t.Fatal("Error", Cos(1.0))
	}
}
