# go-sprite32x32

Package **sprite32x32** provides a type that represents a **32×32 sprite**,
that seamlessly works with Go's built-in `"image"`, `"image/color"`, and `"image/draw"` packages.


## Documention

Online documentation, which includes examples, can be found at: http://godoc.org/github.com/reiver/go-sprite32x32

[![GoDoc](https://godoc.org/github.com/reiver/go-sprite32x32?status.svg)](https://godoc.org/github.com/reiver/go-sprite32x32)


## Example
```go
import (
	"github.com/reiver/go-sprite32x32"
)

// ...

sprite := sprite32x32.Paletted{
	Pix:      pix,
	Palette:  palette,
	Category: "tiles",
	ID:       5,
}

// ...

draw.Draw(dst, rect, sprite, image.ZP, draw.Src)
```
