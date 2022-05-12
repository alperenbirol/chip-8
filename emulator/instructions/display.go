package instructions

import (
	"github.com/alperenbirol/chip-8/emuconfig"
	"github.com/alperenbirol/chip-8/emulator/vm"
)

func Display(xPosRegister, yPosRegister, size byte) Instruction {
	return func(vm *vm.VirtualMachine) {
		x, y := byte(vm.Registers[xPosRegister])%emuconfig.DISPLAY_WIDTH, byte(vm.Registers[yPosRegister])%emuconfig.DISPLAY_HEIGHT

		vm.Registers[0xF] = 0

		bytes := vm.RAM.FetchBytes(vm.IndexRegister.Get(), uint16(size))

		if vm.Display.Draw(x, y, bytes) {
			vm.Registers[0xF] = 1
		}

		vm.IsDrawing = true
	}
}
