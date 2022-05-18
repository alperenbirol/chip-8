package instructions

import (
	"github.com/alperenbirol/chip-8/emuconfig"
	"github.com/alperenbirol/chip-8/emulator/vm"
)

func SkipIfKeyPressed(registerIndex emuconfig.Nibble) Instruction {
	return func(vm *vm.VirtualMachine) {
		keyIndex := vm.Registers[registerIndex].Get()
		if vm.Keypad.IsKeyPressed(keyIndex) {
			vm.PC.NextInstruction()
		}
	}
}

func SkipIfKeyNotPressed(registerIndex emuconfig.Nibble) Instruction {
	return func(vm *vm.VirtualMachine) {
		keyIndex := vm.Registers[registerIndex].Get()
		if !vm.Keypad.IsKeyPressed(keyIndex) {
			vm.PC.NextInstruction()
		}
	}
}

func WaitForKeyPress(registerIndex emuconfig.Nibble) Instruction {
	return func(vm *vm.VirtualMachine) {
		aKeyIsPressed, pressedKey := vm.Keypad.AnyKeysPressed()
		if aKeyIsPressed {
			vm.Registers[registerIndex].Set(pressedKey)
		} else {
			vm.PC.Decrement()
		}
	}
}
