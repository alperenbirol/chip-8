package instructions

import (
	"github.com/alperenbirol/chip-8/emuconfig"
	"github.com/alperenbirol/chip-8/emulator/vm"
)

func SetDelayTimer(registerIndex emuconfig.Nibble) Instruction {
	return func(vm *vm.VirtualMachine) {
		registerValue := vm.Registers[registerIndex].Get()
		vm.DelayTimer.SetTimer(registerValue)
	}
}
