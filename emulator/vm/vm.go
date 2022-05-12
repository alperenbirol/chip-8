package vm

import (
	"github.com/alperenbirol/chip-8/emulator/beeper"
	"github.com/alperenbirol/chip-8/emulator/display"
	"github.com/alperenbirol/chip-8/emulator/indexregister"
	"github.com/alperenbirol/chip-8/emulator/memory"
	"github.com/alperenbirol/chip-8/emulator/programcounter"
	"github.com/alperenbirol/chip-8/emulator/programregister"
	"github.com/alperenbirol/chip-8/emulator/timer"
)

type VirtualMachine struct {
	Display       display.IDisplay
	PC            programcounter.IProgramCounter
	RAM           memory.IMemory
	DelayTimer    timer.IDelayTimer
	SoundTimer    timer.ISoundTimer
	Registers     [16]programregister.ProgramRegister
	IndexRegister indexregister.IndexRegister

	IsDrawing bool
}

func NewVirtualMachine(beeper beeper.IBeeper) *VirtualMachine {
	return &VirtualMachine{
		Display:       display.NewDisplay(),
		PC:            programcounter.NewProgramCounter(),
		RAM:           memory.NewMemory(),
		DelayTimer:    timer.NewDelayTimer(),
		SoundTimer:    timer.NewSoundTimer(beeper),
		Registers:     [16]programregister.ProgramRegister{},
		IndexRegister: 0x00,

		IsDrawing: false,
	}
}
