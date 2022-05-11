package memory

import "github.com/alperenbirol/chip-8/emuconfig"

type Memory struct {
	ram emuconfig.Ram
}

func (m *Memory) Clear() {
	for i := 0x200; i < len(m.ram); i++ {
		m.ram[i] = 0x00
	}
}
