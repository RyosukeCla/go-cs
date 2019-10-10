
bits:
	cd ./pkg/bits; go test

huffman:
	cd ./pkg/huffman; go test

lz77:
	cd ./pkg/lz77; go test

runlength:
	cd ./pkg/runlength; go test

adt:
	cd ./pkg/adt/bstree; go test
	cd ./pkg/adt/treap; go test
	cd ./pkg/adt/heap; go test

sample:
	cd ./pkg/sample; go test

.PHONY: bitstream huffman sample
