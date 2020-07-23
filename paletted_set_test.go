package sprite32x32_test

import (
	"github.com/reiver/go-sprite32x32"

	"github.com/reiver/go-palette2048"
	"github.com/reiver/go-palette2048_tia"

	"image/color"

	"testing"
)

func TestPaletted_Set(t *testing.T) {

		var palette palette2048.Slice = palette2048.Slice(palette2048_tia.Palette[:])

		const category = "does not matter category"
		const id = 7

		var pix [32*32*1]uint8
		var sprite sprite32x32.Paletted = sprite32x32.Paletted{
			Pix: pix[:],
			Palette: palette,
			Category: category,
			ID: id,
		}

		{
			for x:=0; x<32; x++ {
				if 0 ==x%2 {
					sprite.Set(x,0, color.RGBA{  0,  0,  0, 255})
				} else {
					sprite.Set(x,0, color.RGBA{255,255,255, 255})
				}
			}

			for x:=0; x<32; x++ {
				sprite.Set(x,1, color.RGBA{255,  0,  0, 255})
				sprite.Set(x,2, color.RGBA{255,165,  0, 255})
				sprite.Set(x,3, color.RGBA{255,255,  0, 255})
				sprite.Set(x,4, color.RGBA{  0,255,  0, 255})
				sprite.Set(x,5, color.RGBA{  0,  0,255, 255})
				sprite.Set(x,6, color.RGBA{ 75,  0,130, 255})
				sprite.Set(x,7, color.RGBA{238,130,238, 255})
			}
			for y:=8; y<32; y++ {
				for x:=0; x<32; x++ {
					sprite.Set(x,y, color.RGBA{uint8(y),  0,uint8(y), 255})
				}
			}
		}

		var indexes [32][32]uint8
		{
			for y:=0; y<32; y++ {
				for x:=0; x<32; x++ {
					indexes[x][y] = sprite.ColorIndexAt(x,y)
				}
			}
		}
		{
			for x:=0; x<32; x++ {
				const y = 0

				actual := indexes[x][y]
				var expected uint8
				if 0==x%2 {
					expected = 0
				} else {
					expected = 7
				}

				if expected != actual {
					t.Errorf("The actual color index at (x,y)=(%d,%d) is not what was expected.", x,y)
					t.Logf("EXPECTED: %d -> %s", expected, palette.Color(expected))
					t.Logf("ACTUAL:   %d -> %s", actual,   palette.Color(actual))
				}
			}

			for x:=0; x<32; x++ {
				const y = 1

				actual := indexes[x][y]
				const expected = 33

				if expected != actual {
					t.Errorf("The actual color index at (x,y)=(%d,%d) is not what was expected.", x,y)
					t.Logf("EXPECTED: %d -> %s", expected, palette.Color(expected))
					t.Logf("ACTUAL:   %d -> %s", actual,   palette.Color(actual))
				}
			}

			for x:=0; x<32; x++ {
				const y = 2

				actual := indexes[x][y]
				const expected = 12

				if expected != actual {
					t.Errorf("The actual color index at (x,y)=(%d,%d) is not what was expected.", x,y)
					t.Logf("EXPECTED: %d -> %s", expected, palette.Color(expected))
					t.Logf("ACTUAL:   %d -> %s", actual,   palette.Color(actual))
				}
			}

			for x:=0; x<32; x++ {
				const y = 3

				actual := indexes[x][y]
				const expected = 14

				if expected != actual {
					t.Errorf("The actual color index at (x,y)=(%d,%d) is not what was expected.", x,y)
					t.Logf("EXPECTED: %d -> %s", expected, palette.Color(expected))
					t.Logf("ACTUAL:   %d -> %s", actual,   palette.Color(actual))
				}
			}

			for x:=0; x<32; x++ {
				const y = 4

				actual := indexes[x][y]
				const expected = 98

				if expected != actual {
					t.Errorf("The actual color index at (x,y)=(%d,%d) is not what was expected.", x,y)
					t.Logf("EXPECTED: %d -> %s", expected, palette.Color(expected))
					t.Logf("ACTUAL:   %d -> %s", actual,   palette.Color(actual))
				}
			}

			for x:=0; x<32; x++ {
				const y = 5

				actual := indexes[x][y]
				const expected = 65

				if expected != actual {
					t.Errorf("The actual color index at (x,y)=(%d,%d) is not what was expected.", x,y)
					t.Logf("EXPECTED: %d -> %s", expected, palette.Color(expected))
					t.Logf("ACTUAL:   %d -> %s", actual,   palette.Color(actual))
				}
			}

			for x:=0; x<32; x++ {
				const y = 6

				actual := indexes[x][y]
				const expected = 48

				if expected != actual {
					t.Errorf("The actual color index at (x,y)=(%d,%d) is not what was expected.", x,y)
					t.Logf("EXPECTED: %d -> %s", expected, palette.Color(expected))
					t.Logf("ACTUAL:   %d -> %s", actual,   palette.Color(actual))
				}
			}

			for x:=0; x<32; x++ {
				const y = 7

				actual := indexes[x][y]
				const expected = 46

				if expected != actual {
					t.Errorf("The actual color index at (x,y)=(%d,%d) is not what was expected.", x,y)
					t.Logf("EXPECTED: %d -> %s", expected, palette.Color(expected))
					t.Logf("ACTUAL:   %d -> %s", actual,   palette.Color(actual))
				}
			}

			for y:=8; y<32; y++ {
				for x:=0; x<32; x++ {
					actual := indexes[x][y]
					const expected = 0

					if expected != actual {
						t.Errorf("The actual color index at (x,y)=(%d,%d) is not what was expected.", x,y)
						t.Logf("EXPECTED: %d -> %s", expected, palette.Color(expected))
						t.Logf("ACTUAL:   %d -> %s", actual,   palette.Color(actual))
					}
				}
			}
		}
}
