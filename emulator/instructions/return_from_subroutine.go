package instructions

import (
	"fmt"

	"github.com/alperenbirol/chip-8/emuconfig"
	"github.com/alperenbirol/chip-8/emulator/vm"
)

func ReturnFromSubroutine() Instruction {
	return func(vm *vm.VirtualMachine) {
		addr, err := vm.Stack.Pop()
		if err != nil {
			panic(fmt.Errorf("error when returning from subroutine: %v", err))
		}
		vm.PC.SetToAddress(emuconfig.Address(addr))
	}
}
