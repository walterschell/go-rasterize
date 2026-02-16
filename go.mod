module github.com/walterschell/go-rasterize

go 1.24.4

require (
	github.com/srwiley/oksvg v0.0.0-20221011165216-be6e8873101c
	github.com/srwiley/rasterx v0.0.0-20220730225603-2ab79fcdd4ef
)

require (
	golang.org/x/image v0.36.0 // indirect
	golang.org/x/net v0.50.0 // indirect
	golang.org/x/text v0.34.0 // indirect
)

replace github.com/srwiley/oksvg => ./internal/oksvg
