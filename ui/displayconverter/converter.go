package displayconverter

import (
	"image"

	"github.com/alperenbirol/chip-8/emuconfig"
)

func Convert(displayBits emuconfig.Pixels) *image.RGBA {
	rgba := image.NewRGBA(emuconfig.IMAGE_BOUNDS)

	for Yindex, row := range displayBits {
		for Xindex, pixel := range row {
			if pixel {
				rgba.Set(Xindex, Yindex, emuconfig.PIXEL_ON_COLOR)
			} else {
				rgba.Set(Xindex, Yindex, emuconfig.PIXEL_OFF_COLOR)
			}
		}
	}

	return rgba
}
