package sprite32x32_test

import (
	"github.com/reiver/go-sprite32x32"

	"image"

	"testing"
)

func TestImageImage(t *testing.T) {

	// THIS IS WHAT ACTUALLY MATTERS!
	var x image.Image = sprite32x32.Paletted{}

	if nil == x {
		t.Error("This should never happen")
	}
}
