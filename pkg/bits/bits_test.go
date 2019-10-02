package bits

import (
	"reflect"
	"testing"
)

func TestFromByte(t *testing.T) {
	actual := FromByte(20)
	expected := []byte{0, 0, 0, 1, 0, 1, 0, 0}

	if !reflect.DeepEqual(expected, actual) {
		t.Fatal("Fail")
	}

	actual = FromByte(255)
	expected = []byte{1, 1, 1, 1, 1, 1, 1, 1}

	if !reflect.DeepEqual(expected, actual) {
		t.Fatal("Fail")
	}
}

func TestToByte(t *testing.T) {
	actual := ToByte([]byte{0, 0, 0, 1, 0, 1, 0, 0})
	expected := byte(20)

	if !reflect.DeepEqual(expected, actual) {
		t.Fatal("Fail")
	}

	actual = ToByte([]byte{1, 1, 1, 1, 1, 1, 1, 1})
	expected = byte(255)

	if !reflect.DeepEqual(expected, actual) {
		t.Fatal("Fail")
	}
}

func TestBitStream_Write(t *testing.T) {
	bits := []uint8{1, 1, 1, 0, 1, 1, 0, 1, 1, 1, 1, 0, 1, 0, 0, 1, 1, 0, 1, 0, 0, 0, 0, 1, 0}
	bs := NewBitStream([]byte{})
	bs.Write(bits)

	if !reflect.DeepEqual(bs.Bits(), bits) {
		t.Fatal("Fail")
	}
}

func TestBitStream_WriteFromBytes(t *testing.T) {
	bytes := []byte{255, 20}
	bits := []byte{1, 1, 1, 1, 1, 1, 1, 1, 0, 0, 0, 1, 0, 1, 0, 0}
	bs := NewBitStream([]byte{})
	bs.WriteFromBytes(bytes)
	if !reflect.DeepEqual(bs.Bits(), bits) {
		t.Fatal("Fail")
	}
}

func TestBitStream_Read(t *testing.T) {
	bs := NewBitStream([]byte{})

	bs.Write([]byte{0, 0, 0, 0, 0, 0, 0, 1})
	bits := make([]byte, 7)
	bs.Read(bits)

	if !reflect.DeepEqual(bits, []byte{0, 0, 0, 0, 0, 0, 0}) {
		t.Fatal("Fail")
	}
	if !reflect.DeepEqual(bs.Bits(), []byte{1}) {
		t.Fatal("Fail")
	}
	if bs.Len() != 1 {
		t.Fatal("Fail")
	}
}

func TestBitStream_ReadAsBytes(t *testing.T) {
	bits := []byte{1, 1, 1, 1, 1, 1, 1, 1, 0, 0, 0, 1, 0, 1, 0, 0} // 255, 20,
	bs := NewBitStream([]byte{})
	bs.Write(bits)
	bytes := make([]byte, 2)
	bs.ReadAsBytes(bytes)

	if !reflect.DeepEqual(bytes, []byte{255, 20}) {
		t.Fatal("Fail")
	}
	if !reflect.DeepEqual(bs.Bits(), []byte{}) {
		t.Fatal("Fail")
	}
	if bs.Len() != 0 {
		t.Fatal("Fail")
	}

	bs = NewBitStream([]byte{})
	bs.Write([]byte{0})
	bytes = make([]byte, 1)
	n, err := bs.ReadAsBytes(bytes)
	if n != 1 || err == nil {
		t.Fatal("Fail")
	}
}
