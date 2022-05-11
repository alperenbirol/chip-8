package memory

import "github.com/alperenbirol/chip-8/emuconfig"

type Memory struct {
	ram emuconfig.Ram
}

func (m *Memory) loadAtAddress(address uint16, data []byte) {
	copy(m.ram[address:], data)
}

func (m *Memory) UnloadROM() {
	for i := 0x200; i < len(m.ram); i++ {
		m.ram[i] = 0x00
	}
}

func (m *Memory) LoadROM(rom []byte) {
	m.loadAtAddress(0x200, rom)
}

func (m *Memory) LoadFonts(fonts []byte) {
	m.loadAtAddress(0x50, fonts)
}
