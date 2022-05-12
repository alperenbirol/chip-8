package instructions

import "github.com/alperenbirol/chip-8/emulator/vm"

func AddToRegister(registerIndex byte, value byte) Instruction {
	return func(vm *vm.VirtualMachine) {
		vm.Registers[registerIndex].Add(value)
	}
}
