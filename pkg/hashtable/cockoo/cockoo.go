package cockoo

import (
	"github.com/RyosukeCla/go-cs/pkg/hash/md5"
	math "github.com/RyosukeCla/go-cs/pkg/math/basic"
)

type Cockoo struct {
	t1, t2  []*entry
	size    int
	maxLoop int
}

type entry struct {
	key   string
	value interface{}
}

func New(size int) *Cockoo {
	t1 := make([]*entry, size, size)
	t2 := make([]*entry, size, size)
	maxLoop := int(math.Log(1+0.05, float64(size)))
	return &Cockoo{
		t1,
		t2,
		size,
		maxLoop,
	}
}

func (cf *Cockoo) hash(key []byte) (int, int) {
	uints := md5.Hash(key)
	return int(uints[0]) % cf.size, int(uints[1]) % cf.size
}

func (cf *Cockoo) Insert(key string, value interface{}) {
	if cf.Lookup(key) {
		return
	}

	x := &entry{
		key,
		value,
	}
	for i := 0; i < cf.maxLoop; i++ {
		h1, h2 := cf.hash([]byte(x.key))

		// swap x <-> t1[h1(x)]
		temp := cf.t1[h1]
		cf.t1[h1] = x
		x = temp
		if x == nil {
			return
		}

		// swap x <-> t2[h2(x)]
		temp = cf.t2[h2]
		cf.t1[h1] = x
		x = temp
		if x == nil {
			return
		}
	}

	cf.Rehash()
	cf.Insert(x.key, x.value)
}

func (cf *Cockoo) Lookup(key string) bool {
	h1, h2 := cf.hash([]byte(key))
	entry := cf.t1[h1]
	if entry != nil && entry.key == key {
		return true
	}
	entry = cf.t2[h2]
	if entry != nil && entry.key == key {
		return true
	}
	return false
}

func (cf *Cockoo) Get(key string) interface{} {
	h1, h2 := cf.hash([]byte(key))
	entry := cf.t1[h1]
	if entry != nil && entry.key == key {
		return entry.value
	}
	entry = cf.t2[h2]
	if entry != nil && entry.key == key {
		return entry.value
	}
	return nil
}

func (cf *Cockoo) Rehash() {
}
