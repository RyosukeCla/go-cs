module github.com/RyosukeCla/go-playground

go 1.12

replace (
	github.com/RyosukeCla/go-playground/pkg/adt => ./pkg/adt
	github.com/RyosukeCla/go-playground/pkg/adt/bstree => ./pkg/adt/bstree
	github.com/RyosukeCla/go-playground/pkg/adt/heap => ./pkg/adt/heap
	github.com/RyosukeCla/go-playground/pkg/adt/treap => ./pkg/adt/treap
	github.com/RyosukeCla/go-playground/pkg/bits => ./pkg/bits
	github.com/RyosukeCla/go-playground/pkg/huffman => ./pkg/huffman
	github.com/RyosukeCla/go-playground/pkg/lz77 => ./pkg/lz77
	github.com/RyosukeCla/go-playground/pkg/runlength => ./pkg/runlength
	github.com/RyosukeCla/go-playground/pkg/sample => ./pkg/sample
)

require (
	github.com/RyosukeCla/go-playground/pkg/adt v0.0.0
	github.com/RyosukeCla/go-playground/pkg/adt/bstree v0.0.0
	github.com/RyosukeCla/go-playground/pkg/adt/heap v0.0.0
	github.com/RyosukeCla/go-playground/pkg/adt/treap v0.0.0
	github.com/RyosukeCla/go-playground/pkg/bits v0.0.0
	github.com/RyosukeCla/go-playground/pkg/huffman v0.0.0
	github.com/RyosukeCla/go-playground/pkg/lz77 v0.0.0
	github.com/RyosukeCla/go-playground/pkg/runlength v0.0.0
	github.com/RyosukeCla/go-playground/pkg/sample v0.0.0
)
