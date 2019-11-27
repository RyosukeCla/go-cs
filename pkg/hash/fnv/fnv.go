package fnv

// BASIS ...
const BASIS = 2166136261

// PRIME ...
const PRIME = 1099511628211

// Hash ...
func Hash(bytes []byte) int {
	hash := BASIS
	for b := range bytes {
		hash = hash ^ b // xor
		hash = hash * PRIME
	}
	return hash
}
