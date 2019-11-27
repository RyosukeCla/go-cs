test-all:
	make test-bits
	make test-coding
	make test-adt
	make test-sort
	make test-search
	make test-math
	make test-hashtable
	make test-hash
	make test-rand

bench-all:
	make bench-sort

compare-all:
	make compare-quick-with-merge

# tests
test-bits:
	cd ./pkg/bits; go test

test-coding:
	cd ./pkg/coding/huffman; go test
	cd ./pkg/coding/lz77; go test
	cd ./pkg/coding/runlength; go test

test-adt:
	cd ./pkg/adt/bstree; go test
	cd ./pkg/adt/treap; go test
	cd ./pkg/adt/heap; go test

test-sort:
	cd ./pkg/sort/bubble; go test
	cd ./pkg/sort/bucket; go test
	cd ./pkg/sort/heap; go test
	cd ./pkg/sort/merge; go test
	cd ./pkg/sort/quick; go test
	cd ./pkg/sort/insertion; go test
	cd ./pkg/sort/intro; go test

test-search:
	cd ./pkg/search/linear; go test;
	cd ./pkg/search/binary; go test;
	cd ./pkg/search/jump; go test;

test-math:
	cd ./pkg/math/basic; go test;
	cd ./pkg/math/newton; go test;

test-hashtable:
	cd ./pkg/hashtable/basic; go test;

test-hash:
	cd ./pkg/hash/fnv; go test;

test-rand:
	cd ./pkg/rand/midsquare; go test;

# benches
bench-sort:
	cd ./pkg/sort; go test -bench . -benchmem

# comparisons
compare-quick-with-merge:
	cd ./pkg/sort; go test

# etc
sample:
	cd ./pkg/sample; go test

