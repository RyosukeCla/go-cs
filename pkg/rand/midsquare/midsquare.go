package midsquare

import (
	"fmt"
	"strconv"

	"github.com/RyosukeCla/go-cs/pkg/hashtable/basic"
)

// Generate generates pseudorandom. seed is required to be 4 digits number.
func Generate(seed int) int {
	hashtable := basic.NewHashTable(100000)
	digits := seed
	for {
		digitsStr := fmt.Sprintf("%d", digits)
		if hashtable.Get(digitsStr) == true {
			break
		}

		hashtable.Put(string(digits), true)
		squared := digits * digits
		str := fmt.Sprintf("%d", squared)
		digits, _ = strconv.Atoi(str[2:6])
	}

	return digits
}
