package sprite32x32

import (
	"image/color"
)

type Palette interface {
	color.Model
	Color(uint8) color.Color
}
