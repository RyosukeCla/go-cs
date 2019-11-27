package fnv

// BASIS ...
const BASIS uint = 2166136261

// PRIME ...
const PRIME uint = 16777619

// Hash: fnv-1a implementation
func Hash(bytes []byte) uint {
	hash := BASIS
	for b := range bytes {
		hash = hash * PRIME
		hash = hash ^ uint(b) // xor
	}
	return hash
}
