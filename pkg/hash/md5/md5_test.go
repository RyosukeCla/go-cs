package md5

import (
	"fmt"
	"testing"
)

func TestEncode(t *testing.T) {
	fmt.Println(Hash([]byte("md5 message digest algorith")))
	fmt.Println(Hash([]byte("md5 message eigest algorith")))
}
