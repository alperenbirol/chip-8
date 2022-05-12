package programcounter

import "github.com/alperenbirol/chip-8/emuconfig"

type pc uint16

type IProgramCounter interface {
	NextInstruction()
	Reset()
	Decrement()
	SetToAddress(address emuconfig.Address)
	Get() emuconfig.Address
}

func NewProgramCounter() IProgramCounter {
	pc := pc(0x200)
	return &pc
}

func (p *pc) NextInstruction() {
	*p += 2
}

func (p *pc) Decrement() {
	*p -= 2
}

func (p *pc) Reset() {
	*p = 0x200
}

func (p *pc) SetToAddress(address emuconfig.Address) {
	*p = pc(address)
}

func (p *pc) Get() emuconfig.Address {
	return emuconfig.Address(*p)
}
