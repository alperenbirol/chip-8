package vm

import (
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
}
