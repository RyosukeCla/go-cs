
bits:
	cd ./pkg/bits; go test

coding:
	cd ./pkg/coding/huffman; go test
	cd ./pkg/coding/lz77; go test
	cd ./pkg/coding/runlength; go test

adt:
	cd ./pkg/adt/bstree; go test
	cd ./pkg/adt/treap; go test
	cd ./pkg/adt/heap; go test

sort:
	cd ./pkg/sort/bubble; go test
	cd ./pkg/sort/bucket; go test
	cd ./pkg/sort/heap; go test
	cd ./pkg/sort/merge; go test
	cd ./pkg/sort/quick; go test
	cd ./pkg/sort/insertion; go test
	cd ./pkg/sort/intro; go test

bench-sort:
	cd ./pkg/sort; go test -bench . -benchmem

compare-quick-with-merge:
	cd ./pkg/sort; go test

search:
	cd ./pkg/search/linear; go test;
	cd ./pkg/search/binary; go test;
	cd ./pkg/search/jump; go test;

math:
	cd ./pkg/math/basic; go test;
	cd ./pkg/math/newton; go test;

hashtable:
	cd ./pkg/hashtable/basic; go test;

hash:
	cd ./pkg/hash/fnv; go test;


sample:
	cd ./pkg/sample; go test

.PHONY: bitstream huffman sample
