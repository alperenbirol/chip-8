package memory

import (
	"github.com/alperenbirol/chip-8/emuconfig"
)

type ram [emuconfig.RAM_SIZE]byte
type Memory struct {
	ram
}

type IMemory interface {
	LoadROM(rom []byte)
	LoadFonts(fonts []byte)
	FetchInstruction(address emuconfig.Address) emuconfig.Opcode
	UnloadROM()
	FetchByte(address emuconfig.Address) byte
	FetchBytes(address emuconfig.Address, length byte) []byte
}

func NewMemory() IMemory {
	return &Memory{}
}

func (m *Memory) loadAtAddress(address emuconfig.Address, data []byte) {
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

func (m *Memory) FetchInstruction(address emuconfig.Address) emuconfig.Opcode {
	return [2]byte{
		m.ram[address], m.ram[address+1],
	}
}

func (m *Memory) FetchByte(address emuconfig.Address) byte {
	return m.ram[address]
}

func (m *Memory) FetchBytes(address emuconfig.Address, length byte) []byte {
	return m.ram[address : address+emuconfig.Address(length)]
}
