package emulator

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
			e.vm.Keypad.PressKey(0x0)
		},
		KEY_1_PRESSED: func() {
			e.vm.Keypad.PressKey(0x1)
		},
		KEY_2_PRESSED: func() {
			e.vm.Keypad.PressKey(0x2)
		},
		KEY_3_PRESSED: func() {
			e.vm.Keypad.PressKey(0x3)
		},
		KEY_4_PRESSED: func() {
			e.vm.Keypad.PressKey(0x4)
		},
		KEY_5_PRESSED: func() {
			e.vm.Keypad.PressKey(0x5)
		},
		KEY_6_PRESSED: func() {
			e.vm.Keypad.PressKey(0x6)
		},
		KEY_7_PRESSED: func() {
			e.vm.Keypad.PressKey(0x7)
		},
		KEY_8_PRESSED: func() {
			e.vm.Keypad.PressKey(0x8)
		},
		KEY_9_PRESSED: func() {
			e.vm.Keypad.PressKey(0x9)
		},
		KEY_A_PRESSED: func() {
			e.vm.Keypad.PressKey(0xA)
		},
		KEY_B_PRESSED: func() {
			e.vm.Keypad.PressKey(0xB)
		},
		KEY_C_PRESSED: func() {
			e.vm.Keypad.PressKey(0xC)
		},
		KEY_D_PRESSED: func() {
			e.vm.Keypad.PressKey(0xD)
		},
		KEY_E_PRESSED: func() {
			e.vm.Keypad.PressKey(0xE)
		},
		KEY_F_PRESSED: func() {
			e.vm.Keypad.PressKey(0xF)
		},
	}
}
