package consistent

import (
	"github.com/RyosukeCla/go-cs/pkg/adt/treap"
	"github.com/RyosukeCla/go-cs/pkg/hash/fnv"
)

// HashTable ...
type HashTable struct {
	circle treap.Treap[node]
}

type node struct {
	hash uint32
	name string
}

func NewHashTable(nodeNames []string) *HashTable {
	circle := treap.NewTreap(comparator)
	for i := 0; i < len(nodeNames); i++ {
		circle.Insert(node{
			hash: hashFunction(nodeNames[i]),
			name: nodeNames[i],
		})
	}
	return &HashTable{
		circle: circle,
	}
}

func comparator(left node, right node) int {
	if left.hash < right.hash {
		return -1
	} else if left.hash > right.hash {
		return 1
	} else {
		return 0
	}
}

func hashFunction(value string) uint32 {
	bytes := []byte(value)
	return fnv.Hash(bytes)
}

func (h *HashTable) Get(key string) *string {
	hash := hashFunction(key)
	var neighborNode *node
	h.circle.Walk(func(node *treap.Node[node]) int {
		if hash == node.Value.hash {
			neighborNode = &node.Value
			return 0
		} else if hash < node.Value.hash {
			// search left
			neighborNode = &node.Value
			return -1
		} else if node.Value.hash < hash {
			// search right
			return 1
		}
		return 1
	})
	if neighborNode == nil {
		_node := h.circle.Min()
		neighborNode = _node
	}
	if neighborNode == nil {
		return nil
	}
	return &neighborNode.name
}

func (h *HashTable) AddNode(nodeName string) {
	node := node{
		hash: hashFunction(nodeName),
		name: nodeName,
	}
	h.circle.Insert(node)
}

func (h *HashTable) RemoveNode(nodeName string) {
	node := node{
		hash: hashFunction(nodeName),
		name: nodeName,
	}
	h.circle.Remove(node)
}
