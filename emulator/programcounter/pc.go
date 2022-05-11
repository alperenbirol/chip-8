package programcounter

import (
	"github.com/alperenbirol/chip-8/emulator/memory"
)

type pc memory.Address

type IProgramCounter interface {
	NextInstruction()
	Reset()
	SetToAddress(address memory.Address)
}

func NewProgramCounter() IProgramCounter {
	pc := pc(0x200)
	return &pc
}

func (p *pc) NextInstruction() {
	*p += 2
}

func (p *pc) Reset() {
	*p = 0x200
}

func (p *pc) SetToAddress(address memory.Address) {
	*p = pc(address)
}
