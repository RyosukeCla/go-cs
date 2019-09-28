module github.com/RyosukeCla/go-playground

go 1.12

replace (
	github.com/RyosukeCla/go-playground/pkg/bitstream => ./pkg/bitstream
	github.com/RyosukeCla/go-playground/pkg/sample => ./pkg/sample
)

require (
	github.com/RyosukeCla/go-playground/pkg/bitstream v0.0.0
	github.com/RyosukeCla/go-playground/pkg/sample v0.0.0
)
