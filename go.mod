module github.com/RyosukeCla/go-playground

go 1.12

replace (
	github.com/RyosukeCla/go-playground/pkg/bits => ./pkg/bits
	github.com/RyosukeCla/go-playground/pkg/huffman => ./pkg/huffman
	github.com/RyosukeCla/go-playground/pkg/sample => ./pkg/sample
)

require (
	github.com/RyosukeCla/go-playground/pkg/bits v0.0.0
	github.com/RyosukeCla/go-playground/pkg/huffman v0.0.0
	github.com/RyosukeCla/go-playground/pkg/sample v0.0.0
)
