package memory

import (
	"github.com/alperenbirol/chip-8/emuconfig"
)

type Memory struct {
	emuconfig.Ram
}

type IMemory interface {
	LoadROM(rom []byte)
	LoadFonts(fonts []byte)
	FetchInstruction(address emuconfig.Address) emuconfig.Opcode
	UnloadROM()
	FetchByte(address emuconfig.Address) byte
	FetchBytes(address emuconfig.Address, length uint16) []byte
	GetRamPointer() *emuconfig.Ram
	StoreBytesAtAddress(address emuconfig.Address, data byte)
}

func NewMemory() IMemory {
	return &Memory{}
}

func (m *Memory) loadAtAddress(address emuconfig.Address, data []byte) {
	copy(m.Ram[address:], data)
}

func (m *Memory) UnloadROM() {
	for i := 0x200; i < len(m.Ram); i++ {
		m.Ram[i] = 0x00
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
		m.Ram[address], m.Ram[address+1],
	}
}

func (m *Memory) FetchByte(address emuconfig.Address) byte {
	return m.Ram[address]
}

func (m *Memory) FetchBytes(address emuconfig.Address, length uint16) []byte {
	return m.Ram[address : address+emuconfig.Address(length)]
}

func (m *Memory) GetRamPointer() *emuconfig.Ram {
	return &m.Ram
}

func (m *Memory) StoreBytesAtAddress(address emuconfig.Address, data byte) {
	m.Ram[address] = data
}
