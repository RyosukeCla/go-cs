
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

sample:
	cd ./pkg/sample; go test

.PHONY: bitstream huffman sample
