package instructions

import (
	"github.com/alperenbirol/chip-8/emuconfig"
	"github.com/alperenbirol/chip-8/emulator/vm"
)

func SetIndexRegisterToFont(registerIndex emuconfig.Nibble) Instruction {
	return func(vm *vm.VirtualMachine) {
		registerValue := vm.Registers[registerIndex].Get()
		fontIndex := uint16(0x50 + registerValue*0x05)
		vm.IndexRegister.Set(uint16(fontIndex))
	}
}
