
bit_stream:
	cd ./pkg/bit_stream; go test

sample:
	cd ./pkg/sample; go test

.PHONY: bit_stream sample