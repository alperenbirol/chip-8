package instructions

import (
	"github.com/alperenbirol/chip-8/emuconfig"
	"github.com/alperenbirol/chip-8/emulator/vm"
)

func AddToIndexRegister(registerIndex emuconfig.Nibble) Instruction {
	return func(vm *vm.VirtualMachine) {
		vm.IndexRegister.Add(uint16(vm.Registers[registerIndex].Get()))
	}
}
