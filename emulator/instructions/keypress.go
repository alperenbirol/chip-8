package instructions

import (
	"github.com/alperenbirol/chip-8/emuconfig"
	"github.com/alperenbirol/chip-8/emulator/vm"
)

func SkipIfKeyNotPressed(registerIndex emuconfig.Nibble) Instruction {
	return func(vm *vm.VirtualMachine) {
		keyIndex := vm.Registers[registerIndex].Get()
		if !vm.Keypad.IsKeyPressed(keyIndex) {
			vm.PC.NextInstruction()
		}
	}
}
