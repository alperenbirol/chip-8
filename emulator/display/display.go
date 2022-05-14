package display

import (
	"fmt"

	"github.com/alperenbirol/chip-8/emuconfig"
)

type Display struct {
	pixels emuconfig.Pixels
}

type IDisplay interface {
	Clear()
	Draw(x, y byte, sprite []byte) bool
	AsciiPrint()
	GetDisplay() emuconfig.Pixels
}

func NewDisplay() IDisplay {
	d := &Display{}
	return d
}

func (d *Display) setPixel(x, y byte, value emuconfig.Pixel) {
	d.pixels[y%emuconfig.DISPLAY_HEIGHT][x%emuconfig.DISPLAY_WIDTH] = value
}

func (d *Display) getPixel(x, y byte) emuconfig.Pixel {
	return d.pixels[y%emuconfig.DISPLAY_HEIGHT][x%emuconfig.DISPLAY_WIDTH]
}

func (d *Display) Clear() {
	for x := byte(0); x < emuconfig.DISPLAY_WIDTH; x++ {
		for y := byte(0); y < emuconfig.DISPLAY_HEIGHT; y++ {
			d.setPixel(x, y, emuconfig.PIXEL_OFF)
		}
	}
}

func (d *Display) Draw(x, y byte, sprite []byte) bool {
	isOverlap := false
	for i := 0; i < len(sprite); i++ {
		binary := fmt.Sprintf("%08b", sprite[i])
		for j := 0; j < 8; j++ {
			if binary[j] == '1' {
				if d.getPixel(x+byte(j), y+byte(i)) == emuconfig.PIXEL_ON {
					isOverlap = true
					d.setPixel(x+byte(j), y+byte(i), emuconfig.PIXEL_OFF)
				} else {
					d.setPixel(x+byte(j), y+byte(i), emuconfig.PIXEL_ON)
				}
			} else {
				d.setPixel(x+byte(j), y+byte(i), emuconfig.PIXEL_OFF)
			}
		}
	}
	return isOverlap
}

func (d *Display) AsciiPrint() {
	for y := byte(0); y < emuconfig.DISPLAY_HEIGHT; y++ {
		for x := byte(0); x < emuconfig.DISPLAY_WIDTH; x++ {
			if d.getPixel(x, y) == emuconfig.PIXEL_ON {
				fmt.Print("█")
			} else {
				fmt.Print("░")
			}
		}
		fmt.Println()
	}
}

func (d *Display) GetDisplay() emuconfig.Pixels {
	return d.pixels
}
