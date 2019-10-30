module github.com/RyosukeCla/go-cs

go 1.12

replace (
	github.com/RyosukeCla/go-cs/pkg/adt => ./pkg/adt
	github.com/RyosukeCla/go-cs/pkg/adt/bstree => ./pkg/adt/bstree
	github.com/RyosukeCla/go-cs/pkg/adt/heap => ./pkg/adt/heap
	github.com/RyosukeCla/go-cs/pkg/adt/treap => ./pkg/adt/treap
	github.com/RyosukeCla/go-cs/pkg/bits => ./pkg/bits
	github.com/RyosukeCla/go-cs/pkg/coding/huffman => ./pkg/coding/huffman
	github.com/RyosukeCla/go-cs/pkg/coding/lz77 => ./pkg/coding/lz77
	github.com/RyosukeCla/go-cs/pkg/coding/runlength => ./pkg/coding/runlength
	github.com/RyosukeCla/go-cs/pkg/sample => ./pkg/sample
)

require (
	github.com/RyosukeCla/go-cs/pkg/adt v0.0.0
	github.com/RyosukeCla/go-cs/pkg/adt/bstree v0.0.0
	github.com/RyosukeCla/go-cs/pkg/adt/heap v0.0.0
	github.com/RyosukeCla/go-cs/pkg/adt/treap v0.0.0
	github.com/RyosukeCla/go-cs/pkg/bits v0.0.0
	github.com/RyosukeCla/go-cs/pkg/coding/huffman v0.0.0
	github.com/RyosukeCla/go-cs/pkg/coding/lz77 v0.0.0
	github.com/RyosukeCla/go-cs/pkg/coding/runlength v0.0.0
	github.com/RyosukeCla/go-cs/pkg/sample v0.0.0
)
