package instructions

import (
	"github.com/alperenbirol/chip-8/emuconfig"
	"github.com/alperenbirol/chip-8/emulator/vm"
)

func StoreRegistersAtMemory(registerIndex emuconfig.Nibble) Instruction {
	return func(vm *vm.VirtualMachine) {
		address := vm.IndexRegister.Get()
		for i := byte(0); i <= byte(registerIndex); i++ {
			vm.RAM.StoreBytesAtAddress(address, vm.Registers[i].Get())
			address++
		}
	}
}

func FillRegistersFromMemory(registerIndex emuconfig.Nibble) Instruction {
	return func(vm *vm.VirtualMachine) {
		address := vm.IndexRegister.Get()
		for i := byte(0); i <= byte(registerIndex); i++ {
			vm.Registers[i].Set(vm.RAM.FetchByte(address))
			address++
		}
	}
}

func StoreBCD(registerIndex emuconfig.Nibble) Instruction {
	return func(vm *vm.VirtualMachine) {
		address := vm.IndexRegister.Get()
		bcd := vm.Registers[registerIndex].Get()
		vm.RAM.StoreBytesAtAddress(address, bcd/100)
		vm.RAM.StoreBytesAtAddress(address+1, (bcd%100)/10)
		vm.RAM.StoreBytesAtAddress(address+2, bcd%10)
	}
}
