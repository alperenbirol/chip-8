package instructions

import (
	"github.com/alperenbirol/chip-8/emuconfig"
	"github.com/alperenbirol/chip-8/emulator/vm"
)

func SetRegisterToDelayTimer(registerIndex emuconfig.Nibble) Instruction {
	return func(vm *vm.VirtualMachine) {
		timePtr := vm.DelayTimer.GetRemainingTimePointer()

		vm.Registers[registerIndex].Set(*timePtr)
	}
}
