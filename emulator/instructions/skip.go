package instructions

import (
	"github.com/alperenbirol/chip-8/emuconfig"
	"github.com/alperenbirol/chip-8/emulator/vm"
)

func SkipIfRegistersHaveSameValue(registerIndex1 emuconfig.Nibble, registerIndex2 emuconfig.Nibble) Instruction {
	return func(vm *vm.VirtualMachine) {
		if vm.Registers[registerIndex1].Get() == vm.Registers[registerIndex2].Get() {
			vm.PC.NextInstruction()
		}
	}
}

func SkipIfRegistersHaveDifferentValue(registerIndex1 emuconfig.Nibble, registerIndex2 emuconfig.Nibble) Instruction {
	return func(vm *vm.VirtualMachine) {
		if vm.Registers[registerIndex1].Get() != vm.Registers[registerIndex2].Get() {
			vm.PC.NextInstruction()
		}
	}
}
