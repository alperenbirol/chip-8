package instructions

import (
	"github.com/alperenbirol/chip-8/emuconfig"
	"github.com/alperenbirol/chip-8/emulator/vm"
)

func SkipIfNotEqual(registerIndex emuconfig.Nibble, value byte) Instruction {
	return func(vm *vm.VirtualMachine) {
		registerValue := vm.Registers[registerIndex].Get()
		if registerValue != value {
			vm.PC.NextInstruction()
		}
	}
}
