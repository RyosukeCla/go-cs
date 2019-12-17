package consistent

import (
	"fmt"
	"sort"

	"github.com/RyosukeCla/go-cs/pkg/hash/fnv"
	"github.com/RyosukeCla/go-cs/pkg/rand/xorshift"
)

// HashTable ...
type HashTable struct {
	circle circle
	rnd    *xorshift.Rand
}

type circle []node

func (c circle) Len() int {
	return len(c)
}
func (f circle) Less(i, j int) bool {
	return f[i].hash < f[j].hash
}
func (f circle) Swap(i, j int) {
	f[i], f[j] = f[j], f[i]
}

type node struct {
	entries []entry
	hash    uint32
	name    string
}

type entry struct {
	key   string
	value interface{}
}

const MODULO = 4294967295

func NewHashTable(nodeSize, nodeCapacity uint32) *HashTable {
	nodes := make([]node, nodeSize, nodeCapacity)
	rnd := xorshift.NewRand(100)
	for i := 0; i < len(nodes); i++ {
		nodes[i] = node{
			entries: make([]entry, 0, 10),
			hash:    rnd.Generate(),
			name:    fmt.Sprintf("n%d", i),
		}
	}
	sort.Sort(circle(nodes))
	return &HashTable{
		circle: nodes,
		rnd:    rnd,
	}
}

func hashFunction(value string) uint32 {
	bytes := []byte(value)
	return fnv.Hash(bytes)
}

func (h *HashTable) Get(key string) interface{} {
	hash := hashFunction(key) % MODULO
	n := len(h.circle)
	// 1 to n-1 th nodes
	for i := 1; i < n; i++ {
		if h.circle[i-1].hash < hash && hash <= h.circle[i].hash { // find most neighbor node
			for _, entry := range h.circle[i].entries {
				if entry.key == key {
					return entry.value
				}
			}
			return nil
		}
	}
	// 0th node
	for _, entry := range h.circle[0].entries {
		if entry.key == key {
			return entry.value
		}
	}
	return nil
}

func (h *HashTable) Put(key string, value interface{}) {
	hash := hashFunction(key) % MODULO
	n := len(h.circle)
	// 1 to n-1 th nodes
	for i := 1; i < n; i++ {
		if h.circle[i-1].hash < hash && hash <= h.circle[i].hash { // find most neighbor node
			h.circle[i].entries = append(h.circle[i].entries, entry{
				key,
				value,
			})
			return
		}
	}

	// 0th node
	h.circle[0].entries = append(h.circle[0].entries, entry{
		key,
		value,
	})
}

func (h *HashTable) Erase(key string) {
	hash := hashFunction(key) % MODULO
	n := len(h.circle)

	// 1 to n-1 th nodes
	for i := 1; i < n; i++ {
		if h.circle[i-1].hash < hash && hash <= h.circle[i].hash { // find most neighbor node
			for index, entry := range h.circle[i].entries {
				if entry.key == key {
					entries := h.circle[i].entries
					entries[0], entries[index] = entries[index], entries[0] // swap
					h.circle[i].entries = entries[1:]                       // erace
					return
				}
			}
			return
		}
	}

	// 0th node
	for index, entry := range h.circle[0].entries {
		if entry.key == key {
			entries := h.circle[0].entries
			entries[0], entries[index] = entries[index], entries[0] // swap
			h.circle[0].entries = entries[1:]                       // erace
			return
		}
	}
}

func (h *HashTable) List() {
	n := len(h.circle)
	for i := 1; i < n; i++ {
		fmt.Println(h.circle[i].name, h.circle[i].hash, len(h.circle[i].entries))
	}
}

func (h *HashTable) AddNode(nodeName string) {
	node := node{
		entries: make([]entry, 0, 10),
		hash:    h.rnd.Generate(),
		name:    nodeName,
	}

	n := len(h.circle)
	for i := 1; i < n; i++ {
		if h.circle[i-1].hash < node.hash && node.hash <= h.circle[i].hash { // find most neighbor node
			return
		}
	}

	sort.Sort(h.circle)
}

func (h *HashTable) RemoveNode(nodeName string) {

}
