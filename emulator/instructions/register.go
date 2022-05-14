package instructions

import (
	"math/rand"

	"github.com/alperenbirol/chip-8/emuconfig"
	"github.com/alperenbirol/chip-8/emulator/vm"
)

func SetRegisterToRandom(registerIndex emuconfig.Nibble, NN uint16) Instruction {
	return func(vm *vm.VirtualMachine) {
		random := rand.Intn(255)
		vm.Registers[registerIndex].Set(byte(random) & byte(NN))
	}
}
