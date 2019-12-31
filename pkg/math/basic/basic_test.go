package basic

import (
	"testing"
)

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

func TestMod(t *testing.T) {
	if Abs(Mod(5.5, 2.2)-1.1) > 0.00001 {
		t.Fatal("Error", Mod(5.5, 2.2))
	}

	if Abs(Mod(-5.5, 2.2)-1.1) > 0.00001 {
		t.Fatal("Error", Mod(-5.5, 2.2))
	}
}

func TestCeil(t *testing.T) {
	if Ceil(5.5) != 6.0 {
		t.Fatal("Error")
	}

	if Ceil(-5.5) != -5.0 {
		t.Fatal("Error")
	}
}

func TestFloor(t *testing.T) {
	if Floor(5.5) != 5.0 {
		t.Fatal("Error")
	}

	if Floor(-5.5) != -6.0 {
		t.Fatal("Error")
	}
}

func TestSqrt(t *testing.T) {
	if Abs(Sqrt(2)-1.41421356237) > 0.000000001 {
		t.Fatal("Error", Sqrt(2))
	}

	if Abs(Sqrt(99)-9.94987437107) > 0.000000001 {
		t.Fatal("Error", Sqrt(99))
	}
}

func TestPower(t *testing.T) {
	if Abs(Power(3.0, 2.0)-9.0) > 0.0000001 {
		t.Fatal("Error", Power(3.0, 2.0))
	}
}

func TestLn(t *testing.T) {
	if Abs(Ln(10.0)-2.30258509299) > 0.0000001 {
		t.Fatal("Error", Ln(10.0))
	}

	if Abs(Ln(542)-6.29526600144) > 0.0000001 {
		t.Fatal("Error", Ln(542))
	}
}

func TestExp(t *testing.T) {
	if Abs(Exp(1.0)-E) > 0.0000001 {
		t.Fatal("Error", Exp(1.0))
	}

	if Abs(Exp(10.0)-22026.4657948) > 0.0000001 {
		t.Fatal("Error", Exp(10.0))
	}
}

func TestSin(t *testing.T) {
	if Abs(Sin(2.0)-0.90929742682) > 0.0000001 {
		t.Fatal("Error", Sin(2.0))
	}

	if Abs(Sin(100*PI)) > 0.0000001 {
		t.Fatal("Error", Sin(100*PI))
	}
}

func TestCos(t *testing.T) {
	if Abs(Cos(1.0)-0.54030230586) > 0.0000001 {
		t.Fatal("Error", Cos(1.0))
	}
}

func TestTan(t *testing.T) {
	if Abs(Tan(1.0)-1.55740772465) > 0.0000001 {
		t.Fatal("Error", Tan(1.0))
	}
}
