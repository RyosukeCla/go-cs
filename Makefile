
bitstream:
	cd ./pkg/bitstream; go test

huffman:
	cd ./pkg/huffman; go test

sample:
	cd ./pkg/sample; go test

.PHONY: bitstream huffman sample
