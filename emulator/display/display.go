package display

import "github.com/alperenbirol/chip-8/emuconfig"

type Display struct {
	pixel [emuconfig.DISPLAY_WIDTH][emuconfig.DISPLAY_HEIGHT]pixel
}

type IDisplay interface {
	Clear()
	TogglePixel(x, y int)
}

func NewDisplay() IDisplay {
	d := &Display{}
	return d
}

func (d *Display) setPixel(x, y int, value pixel) {
	d.pixel[x][y] = value
}

func (d *Display) getPixel(x, y int) pixel {
	return d.pixel[x][y]
}

func (d *Display) Clear() {
	for x := 0; x < emuconfig.DISPLAY_WIDTH; x++ {
		for y := 0; y < emuconfig.DISPLAY_HEIGHT; y++ {
			d.setPixel(x, y, PIXEL_OFF)
		}
	}
}

func (d *Display) TogglePixel(x, y int) {
	d.setPixel(x, y, !d.getPixel(x, y))
}
