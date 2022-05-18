package keypad

import (
	"time"

	"github.com/alperenbirol/chip-8/emuconfig"
)

type Keypad [16]emuconfig.Key

func NewKeypad() *Keypad {
	keypad := Keypad{
		emuconfig.KEY_0,
		emuconfig.KEY_1,
		emuconfig.KEY_2,
		emuconfig.KEY_3,
		emuconfig.KEY_4,
		emuconfig.KEY_5,
		emuconfig.KEY_6,
		emuconfig.KEY_7,
		emuconfig.KEY_8,
		emuconfig.KEY_9,
		emuconfig.KEY_A,
		emuconfig.KEY_B,
		emuconfig.KEY_C,
		emuconfig.KEY_D,
		emuconfig.KEY_E,
		emuconfig.KEY_F,
	}
	return &keypad
}

func (keypad *Keypad) IsKeyPressed(keyIndex byte) bool {
	return keypad[keyIndex].Pressed
}

func (keypad *Keypad) PressKey(keyIndex byte) {
	keypad[keyIndex].Pressed = true
}

func (keypad *Keypad) ReleaseKey(keyIndex byte) {
	keypad[keyIndex].Pressed = false
}

func (keypad *Keypad) AnyKeysPressed() (bool, byte) {
	for _, key := range keypad {
		if key.Pressed {
			return true, key.KeyCode
		}
	}
	return false, 0
}

func (keypad *Keypad) PressAndReleaseDebugKey(keyIndex byte, releaseDelay uint) {
	keypad.PressKey(keyIndex)
	go func(keyIndex byte, releaseDelay uint) {
		time.Sleep(time.Duration(releaseDelay) * time.Millisecond)
		keypad.ReleaseKey(keyIndex)
	}(keyIndex, releaseDelay)
}
