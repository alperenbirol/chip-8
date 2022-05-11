package instructions

import "github.com/alperenbirol/chip-8/emulator/vm"

func SetIndexRegister(value uint16) Instruction {
	return func(vm *vm.VirtualMachine) {
		vm.IndexRegister.Set(value)
	}
}
