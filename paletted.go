package sprite32x32

import (
	"image"
	"image/color"
)

type Paletted struct {
	// Pix holds the pixel data, as palette indices.
	Pix []uint8

	// Palette is the color palette used to resolve the colors.
	Palette Palette

	// Sheet is the name of the sprite sheet.
	Category string

	// ID is the ID of the sprite in the sprite sheet.
	//
	// The implication here is that there can only be 256 sprites
	// in this sprite category, as a ‘uint8’ only has 256 values —
	// 0 to 255.
	ID uint8
}

func (receiver Paletted) At(x, y int) color.Color {
	if ! (image.Point{X:x,Y:y}).In(receiver.Bounds()) {
		return color.RGBA{0,0,0,0}
	}

	index := receiver.ColorIndexAt(x,y)

	color := receiver.Palette.Color(index)

	return color
}

func (receiver Paletted) Bounds() image.Rectangle {
	const x = 0
	const y = 0

	// [x,x+Width) and [y,y+Height)
	return image.Rectangle{
		Min: image.Point{
			X: x,
			Y: y,
		},
		Max: image.Point{
			X: x+Width,
			Y: y+Height,
		},
	}
}

func (receiver Paletted) ColorIndexAt(x, y int) uint8 {
	offset := receiver.PixOffset(x,y)

	index := receiver.Pix[offset]

	return index
}

func (receiver Paletted) ColorModel() color.Model {
	return receiver.Palette
}

func (receiver Paletted) PixOffset(x int, y int) int {
	return y*(Width*Depth) + x*Depth
}

func (receiver Paletted) Set(x, y int, c color.Color) {
	offset := receiver.PixOffset(x,y)
	if offset < 0 || len(receiver.Pix) <= offset {
		return
	}

	index := receiver.Palette.Index(c)

	receiver.Pix[offset] = index
}

func (receiver Paletted) SpriteCategory() string {
	return receiver.Category
}

func (receiver Paletted) SpriteID() uint8 {
	return receiver.ID
}
