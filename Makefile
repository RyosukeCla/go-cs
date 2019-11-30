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
	go test ./pkg/coding/?*/

test-adt:
	go test ./pkg/adt/?*/

test-sort:
	go test ./pkg/sort/?*/

test-search:
	go test ./pkg/search/?*/

test-math:
	go test ./pkg/math/?*/

test-hashtable:
	go test ./pkg/hashtable/?*/

test-hash:
	go test ./pkg/hash/?*/

test-rand:
	go test ./pkg/rand/?*/

# benches
bench-sort:
	go test ./pkg/sort -bench . -benchmem

# comparisons
compare-quick-with-merge:
	go test ./pkg/sort -v

