package md5

import (
	"fmt"
	"testing"
)

func TestEncode(t *testing.T) {
	fmt.Println(Hash([]byte("message")))
}
