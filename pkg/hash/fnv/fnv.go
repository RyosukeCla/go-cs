package fnv

// BASIS ...
const BASIS uint32 = 2166136261

// PRIME ...
const PRIME uint32 = 16777619

// Hash is fnv-1a implementation
func Hash(bytes []byte) uint32 {
	hash := BASIS
	for _, b := range bytes {
		hash = hash * PRIME
		hash = hash ^ uint32(b) // xor
	}
	return hash
}
