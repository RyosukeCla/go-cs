package fnv

// BASIS ...
const BASIS int = 2166136261

// PRIME ...
const PRIME int = 16777619

// Hash ...
func Hash(bytes []byte) int {
	hash := BASIS
	for b := range bytes {
		hash = hash ^ b // xor
		hash = hash * PRIME
	}
	return hash
}
