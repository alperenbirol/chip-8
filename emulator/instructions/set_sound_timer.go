package instructions

import (
	"github.com/alperenbirol/chip-8/emuconfig"
	"github.com/alperenbirol/chip-8/emulator/vm"
)

func SetSoundTimer(registerIndex emuconfig.Nibble) Instruction {
	return func(vm *vm.VirtualMachine) {
		registerValue := vm.Registers[registerIndex].Get()
		vm.SoundTimer.SetTimer(registerValue)
	}
}
