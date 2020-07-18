package sprite32x32

import (
	"image"
	"image/color"
)

type Paletted struct {
	// Pix holds the pixel data, as palette indices.
	Pix []uint8

	// Palette is the color palette used to resolve the colors.
	Palette color.Palette

	// Sheet is the name of the sprite sheet.
	Category string

	// ID is the ID of the sprite in the sprite sheet.
	ID uint8
}

func (receiver Paletted) At(x, y int) color.Color {
	index := receiver.ColorIndexAt(x,y)

	color := receiver.Palette[index]

	return color
}

func (receiver Paletted) Bounds() image.Rectangle {
	const x = 0
	const y = 0

	// [x,x+width) and [y,y+height)
	return image.Rectangle{
		Min: image.Point{
			X: x,
			Y: y,
		},
		Max: image.Point{
			X: x+width,
			Y: y+width,
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
	return y*(width*depth) + x*depth
}

func (receiver Paletted) SpriteCategory() string {
	return receiver.Category
}

func (receiver Paletted) SpriteID() uint8 {
	return receiver.ID
}
