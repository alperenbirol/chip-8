package emuconfig

import (
	"image"
	"time"

	"image/color"

	"github.com/AllenDang/giu"
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

	DEBUG_WIDGET_FLAGS = giu.WindowFlagsNoCollapse | giu.WindowFlagsNoMove | giu.WindowFlagsNoResize

	DEBUG_KEY_PRESS_DURATION = 100
)

var (
	DEFAULT_FONT_SET = [80]byte{
		0xF0, 0x90, 0x90, 0x90, 0xF0, // 0 0x50
		0x20, 0x60, 0x20, 0x20, 0x70, // 1 0x55
		0xF0, 0x10, 0xF0, 0x80, 0xF0, // 2 0x60
		0xF0, 0x10, 0xF0, 0x10, 0xF0, // 3 0x65
		0x90, 0x90, 0xF0, 0x10, 0x10, // 4 0x70
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

	KEY_0 = Key{0x0, false}
	KEY_1 = Key{0x1, false}
	KEY_2 = Key{0x2, false}
	KEY_3 = Key{0x3, false}
	KEY_4 = Key{0x4, false}
	KEY_5 = Key{0x5, false}
	KEY_6 = Key{0x6, false}
	KEY_7 = Key{0x7, false}
	KEY_8 = Key{0x8, false}
	KEY_9 = Key{0x9, false}
	KEY_A = Key{0xA, false}
	KEY_B = Key{0xB, false}
	KEY_C = Key{0xC, false}
	KEY_D = Key{0xD, false}
	KEY_E = Key{0xE, false}
	KEY_F = Key{0xF, false}

	USER_KEYPAD_INTERFACE = map[giu.Key]Key{
		giu.Key1: KEY_1,
		giu.Key2: KEY_2,
		giu.Key3: KEY_3,
		giu.Key4: KEY_C,
		giu.KeyQ: KEY_4,
		giu.KeyW: KEY_5,
		giu.KeyE: KEY_6,
		giu.KeyR: KEY_D,
		giu.KeyA: KEY_7,
		giu.KeyS: KEY_8,
		giu.KeyD: KEY_9,
		giu.KeyF: KEY_E,
		giu.KeyZ: KEY_A,
		giu.KeyX: KEY_0,
		giu.KeyC: KEY_B,
		giu.KeyV: KEY_F,
	}
	REVERSE_USER_KEYPAD_INTERFACE = map[Key]giu.Key{
		KEY_1: giu.Key1,
		KEY_2: giu.Key2,
		KEY_3: giu.Key3,
		KEY_C: giu.Key4,
		KEY_4: giu.KeyQ,
		KEY_5: giu.KeyW,
		KEY_6: giu.KeyE,
		KEY_D: giu.KeyR,
		KEY_7: giu.KeyA,
		KEY_8: giu.KeyS,
		KEY_9: giu.KeyD,
		KEY_E: giu.KeyF,
		KEY_A: giu.KeyZ,
		KEY_0: giu.KeyX,
		KEY_B: giu.KeyC,
		KEY_F: giu.KeyV,
	}
)
