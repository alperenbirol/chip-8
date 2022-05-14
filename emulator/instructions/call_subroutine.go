package instructions

import (
	"github.com/alperenbirol/chip-8/emuconfig"
	"github.com/alperenbirol/chip-8/emulator/vm"
)

func CallSubroutine(address emuconfig.Address) Instruction {
	return func(vm *vm.VirtualMachine) {
		vm.Stack.Push(uint16(vm.PC.Get()))
		vm.PC.SetToAddress(address)
	}
}
