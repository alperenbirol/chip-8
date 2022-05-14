package instructions

import (
	"github.com/alperenbirol/chip-8/emuconfig"
	"github.com/alperenbirol/chip-8/emulator/vm"
)

func SetRegisterToRegister(registerIndex1 emuconfig.Nibble, registerIndex2 emuconfig.Nibble) Instruction {
	return func(vm *vm.VirtualMachine) {
		vm.Registers[registerIndex1].Set(vm.Registers[registerIndex2].Get())
	}
}

func SetRegisterToRegisterOr(registerIndex1 emuconfig.Nibble, registerIndex2 emuconfig.Nibble) Instruction {
	return func(vm *vm.VirtualMachine) {
		vm.Registers[registerIndex1].Set(vm.Registers[registerIndex1].Get() | vm.Registers[registerIndex2].Get())
	}
}

func SetRegisterToRegisterAnd(registerIndex1 emuconfig.Nibble, registerIndex2 emuconfig.Nibble) Instruction {
	return func(vm *vm.VirtualMachine) {
		vm.Registers[registerIndex1].Set(vm.Registers[registerIndex1].Get() & vm.Registers[registerIndex2].Get())
	}
}

func SetRegisterToRegisterXor(registerIndex1 emuconfig.Nibble, registerIndex2 emuconfig.Nibble) Instruction {
	return func(vm *vm.VirtualMachine) {
		vm.Registers[registerIndex1].Set(vm.Registers[registerIndex1].Get() ^ vm.Registers[registerIndex2].Get())
	}
}

func AddRegisterToRegister(registerIndex1 emuconfig.Nibble, registerIndex2 emuconfig.Nibble) Instruction {
	return func(vm *vm.VirtualMachine) {
		// set carry to vm.Registers[15]
		total := uint16(vm.Registers[registerIndex1].Get()) + uint16(vm.Registers[registerIndex2].Get())
		if total > 255 {
			vm.Registers[15].Set(1)
		} else {
			vm.Registers[15].Set(0)
		}
		vm.Registers[registerIndex1].Set(byte(total))
	}
}

func SubtractRegisterFromRegister(registerIndex1 emuconfig.Nibble, registerIndex2 emuconfig.Nibble) Instruction {
	return func(vm *vm.VirtualMachine) {
		reg1, reg2 := vm.Registers[registerIndex1].Get(), vm.Registers[registerIndex2].Get()
		if reg1 > reg2 {
			vm.Registers[15].Set(1)
		} else {
			vm.Registers[15].Set(0)
		}

		vm.Registers[registerIndex1].Set(reg1 - reg2)
	}
}

func ShiftRegisterRight(registerIndex emuconfig.Nibble, registerIndex2 emuconfig.Nibble) Instruction {
	return func(vm *vm.VirtualMachine) {
		reg1 := vm.Registers[registerIndex].Get()
		vm.Registers[15].Set(reg1 & 1)
		vm.Registers[registerIndex].Set(reg1 >> 1)
	}
}

func ShiftRegisterLeft(registerIndex emuconfig.Nibble, registerIndex2 emuconfig.Nibble) Instruction {
	return func(vm *vm.VirtualMachine) {
		reg1 := vm.Registers[registerIndex].Get()
		vm.Registers[15].Set((reg1 & 0x80) / 0x80)
		vm.Registers[registerIndex].Set(reg1 << 1)
	}
}

func SubtractRegisterFromRegisterReverse(registerIndex1 emuconfig.Nibble, registerIndex2 emuconfig.Nibble) Instruction {
	return func(vm *vm.VirtualMachine) {
		reg1, reg2 := vm.Registers[registerIndex1].Get(), vm.Registers[registerIndex2].Get()
		if reg2 > reg1 {
			vm.Registers[15].Set(1)
		} else {
			vm.Registers[15].Set(0)
		}

		vm.Registers[registerIndex1].Set(reg2 - reg1)
	}
}
