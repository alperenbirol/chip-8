package emuconfig

import (
	"image"
	"time"

	"image/color"
)

const (
	RAM_SIZE = 4096

	DISPLAY_WIDTH  = 64
	DISPLAY_HEIGHT = 32

	TIMER_FREQUENCY = 60
	TIMER_INTERVAL  = time.Second / TIMER_FREQUENCY

	BEEPER_SAMPLE_RATE = 22050
	BEEPER_BUFFER_SIZE = 2048
	BEEPER_FREQUENCY   = 440

	PIXEL_ON  = Pixel(true)
	PIXEL_OFF = Pixel(false)
)

var (
	DEFAULT_FONT_SET = [80]byte{
		0xF0, 0x90, 0x90, 0x90, 0xF0, // 0
		0x20, 0x60, 0x20, 0x20, 0x70, // 1
		0xF0, 0x10, 0xF0, 0x80, 0xF0, // 2
		0xF0, 0x10, 0xF0, 0x10, 0xF0, // 3
		0x90, 0x90, 0xF0, 0x10, 0x10, // 4
		0xF0, 0x80, 0xF0, 0x10, 0xF0, // 5
		0xF0, 0x80, 0xF0, 0x90, 0xF0, // 6
		0xF0, 0x10, 0x20, 0x40, 0x40, // 7
		0xF0, 0x90, 0xF0, 0x90, 0xF0, // 8
		0xF0, 0x90, 0xF0, 0x10, 0xF0, // 9
		0xF0, 0x90, 0xF0, 0x90, 0x90, // A
		0xE0, 0x90, 0xe0, 0x90, 0xE0, // B
		0xF0, 0x80, 0x80, 0x80, 0x80, // C
		0xF0, 0x90, 0x90, 0x90, 0xE0, // D
		0xF0, 0x80, 0xF0, 0x80, 0xF0, // E
		0xF0, 0x80, 0xF0, 0x80, 0x80, // F
	}

	IMAGE_BOUNDS = image.Rect(0, 0, DISPLAY_WIDTH, DISPLAY_HEIGHT)

	PIXEL_ON_COLOR  = color.RGBA{0xff, 0xff, 0xff, 0xff}
	PIXEL_OFF_COLOR = color.RGBA{0x00, 0x00, 0x00, 0xff}
)
