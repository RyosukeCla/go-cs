
bitstream:
	cd ./pkg/bitstream; go test

sample:
	cd ./pkg/sample; go test

.PHONY: bitstream sample