package complex

import (
	"testing"

	math "github.com/RyosukeCla/go-cs/pkg/math/basic"
)

func TestEuler(t *testing.T) {

}

func TestAbs(t *testing.T) {
	if math.AlmostEqual(Abs(New(3, 4)), 5.0) != true {
		t.Fatal("error")
	}
}

func TestArg(t *testing.T) {

}

func TestAdd(t *testing.T) {
	if Equal(Add(New(1, 1), New(2, 2)), New(3, 3)) != true {
		t.Fatal("error")
	}
}

func TestSub(t *testing.T) {
	if Equal(Sub(New(1, 1), New(2, 2)), New(-1, -1)) != true {
		t.Fatal("error")
	}
}

func TestMul(t *testing.T) {

}

func TestReciprocal(t *testing.T) {

}

func TestDiv(t *testing.T) {
}

func TestConjugate(t *testing.T) {
	conj := Conjugate(New(1, 1))
	if Equal(conj, New(1, -1)) != true {
		t.Fatal("error")
	}
}

func TestEqual(t *testing.T) {
	if Equal(New(1, 1), New(1, 1)) != true {
		t.Fatal("error")
	}

	if Equal(New(1, 1), New(-1, -1)) != false {
		t.Fatal("error")
	}
}
