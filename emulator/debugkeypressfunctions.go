package emulator

import "github.com/alperenbirol/chip-8/emuconfig"

type KeypressFunctions struct {
	KEY_0_PRESSED func()
	KEY_1_PRESSED func()
	KEY_2_PRESSED func()
	KEY_3_PRESSED func()
	KEY_4_PRESSED func()
	KEY_5_PRESSED func()
	KEY_6_PRESSED func()
	KEY_7_PRESSED func()
	KEY_8_PRESSED func()
	KEY_9_PRESSED func()
	KEY_A_PRESSED func()
	KEY_B_PRESSED func()
	KEY_C_PRESSED func()
	KEY_D_PRESSED func()
	KEY_E_PRESSED func()
	KEY_F_PRESSED func()
}

func (e *Emulator) KeypressFunctions() *KeypressFunctions {
	return &KeypressFunctions{
		KEY_0_PRESSED: func() {
			e.vm.Keypad.PressAndReleaseDebugKey(0x0, emuconfig.DEBUG_KEY_PRESS_DURATION)
		},
		KEY_1_PRESSED: func() {
			e.vm.Keypad.PressAndReleaseDebugKey(0x1, emuconfig.DEBUG_KEY_PRESS_DURATION)
		},
		KEY_2_PRESSED: func() {
			e.vm.Keypad.PressAndReleaseDebugKey(0x2, emuconfig.DEBUG_KEY_PRESS_DURATION)
		},
		KEY_3_PRESSED: func() {
			e.vm.Keypad.PressAndReleaseDebugKey(0x3, emuconfig.DEBUG_KEY_PRESS_DURATION)
		},
		KEY_4_PRESSED: func() {
			e.vm.Keypad.PressAndReleaseDebugKey(0x4, emuconfig.DEBUG_KEY_PRESS_DURATION)
		},
		KEY_5_PRESSED: func() {
			e.vm.Keypad.PressAndReleaseDebugKey(0x5, emuconfig.DEBUG_KEY_PRESS_DURATION)
		},
		KEY_6_PRESSED: func() {
			e.vm.Keypad.PressAndReleaseDebugKey(0x6, emuconfig.DEBUG_KEY_PRESS_DURATION)
		},
		KEY_7_PRESSED: func() {
			e.vm.Keypad.PressAndReleaseDebugKey(0x7, emuconfig.DEBUG_KEY_PRESS_DURATION)
		},
		KEY_8_PRESSED: func() {
			e.vm.Keypad.PressAndReleaseDebugKey(0x8, emuconfig.DEBUG_KEY_PRESS_DURATION)
		},
		KEY_9_PRESSED: func() {
			e.vm.Keypad.PressAndReleaseDebugKey(0x9, emuconfig.DEBUG_KEY_PRESS_DURATION)
		},
		KEY_A_PRESSED: func() {
			e.vm.Keypad.PressAndReleaseDebugKey(0xA, emuconfig.DEBUG_KEY_PRESS_DURATION)
		},
		KEY_B_PRESSED: func() {
			e.vm.Keypad.PressAndReleaseDebugKey(0xB, emuconfig.DEBUG_KEY_PRESS_DURATION)
		},
		KEY_C_PRESSED: func() {
			e.vm.Keypad.PressAndReleaseDebugKey(0xC, emuconfig.DEBUG_KEY_PRESS_DURATION)
		},
		KEY_D_PRESSED: func() {
			e.vm.Keypad.PressAndReleaseDebugKey(0xD, emuconfig.DEBUG_KEY_PRESS_DURATION)
		},
		KEY_E_PRESSED: func() {
			e.vm.Keypad.PressAndReleaseDebugKey(0xE, emuconfig.DEBUG_KEY_PRESS_DURATION)
		},
		KEY_F_PRESSED: func() {
			e.vm.Keypad.PressAndReleaseDebugKey(0xF, emuconfig.DEBUG_KEY_PRESS_DURATION)
		},
	}
}
