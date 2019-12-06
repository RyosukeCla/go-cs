package base64

import (
	"testing"
)

func TestEncode(t *testing.T) {
	encoded := Encode([]byte("Hello"))
	if string(encoded) != "SGVsbG8=" {
		t.Fatal("error. actual:", string(encoded), ", expeceted: SGVsbG8")
	}
}

func TestDecode(t *testing.T) {
	decoded := Decode([]byte("SGVsbG8="))
	if string(decoded) != "Hello" {
		t.Fatal("error. actual:", string(decoded), ", expeceted: Hello")
	}
}
