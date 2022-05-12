package instructions

import (
	"github.com/alperenbirol/chip-8/emuconfig"
	"github.com/alperenbirol/chip-8/emulator/vm"
)

func Jump(address emuconfig.Address) Instruction {
	return func(vm *vm.VirtualMachine) {
		vm.PC.SetToAddress(address)

		// since pc is incremented after the instruction is executed, we need to decrement it here
		vm.PC.Decrement()
	}
}
