package display

import "github.com/alperenbirol/chip-8/emuconfig"

type Display struct {
	pixel [emuconfig.DISPLAY_WIDTH][emuconfig.DISPLAY_HEIGHT]pixel
}

func (d *Display) SetPixel(x, y int, value pixel) {
	d.pixel[x][y] = value
}

func (d *Display) Clear() {
	for x := 0; x < emuconfig.DISPLAY_WIDTH; x++ {
		for y := 0; y < emuconfig.DISPLAY_HEIGHT; y++ {
			d.SetPixel(x, y, PIXEL_OFF)
		}
	}
}
