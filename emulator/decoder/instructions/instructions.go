package instructions

import "github.com/alperenbirol/chip-8/emulator/vm"

type Instruction func(vm *vm.VirtualMachine)
