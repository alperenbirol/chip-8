package instructions

import "github.com/alperenbirol/chip-8/emulator/vm"

func ClearScreen() Instruction {
	return func(vm *vm.VirtualMachine) {
		vm.Display.Clear()
	}
}
