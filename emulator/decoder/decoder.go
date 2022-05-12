package decoder

import (
	"github.com/alperenbirol/chip-8/emuconfig"
	"github.com/alperenbirol/chip-8/emulator/instructions"
)

func Decode(code *emuconfig.Opcode) instructions.Instruction {
	nibbles := code.GetNibbles()
	switch nibbles[0] {
	case 0x0:
		switch nibbles[1] {
		case 0x0:
			switch nibbles[2] {
			case 0xE:
				switch nibbles[3] {
				case 0x0: // 00E0 - Clears the screen
					return instructions.ClearScreen()
				}
			}
		}
	case 0x1: // 1NNN - Jumps program counter to address NNN
		// convert 3 nibbles to a 16 bit address
		address := uint16(nibbles[1])<<8 | uint16(nibbles[2])<<4 | uint16(nibbles[3])
		return instructions.Jump(emuconfig.Address(address))
	case 0x6: // 6XNN - Sets register X to NN
		// convert 2 nibbles to a byte value
		byteValue := byte(nibbles[2]<<4 | nibbles[3])
		registerIndex := byte(nibbles[1])
		// set VX(registerIndex) register to NN(byteValue)
		return instructions.SetRegister(registerIndex, byteValue)
	case 0x7: // 7XNN - Adds NN to register X
		// convert 2 nibbles to a byte value
		byteValue := byte(nibbles[2]<<4 | nibbles[3])
		registerIndex := byte(nibbles[1])
		// Add NN(byteValue) to VX(registerIndex)
		return instructions.AddToRegister(registerIndex, byteValue)
	case 0xA: // ANNN - Sets index register to NNN
		// convert 3 nibbles to a 16 bit address
		NNN := uint16(nibbles[1])<<8 | uint16(nibbles[2])<<4 | uint16(nibbles[3])
		return instructions.SetIndexRegister(NNN)
	case 0xD: // DXYN - Draws a sprite at coordinate (X,Y) with N bytes of sprite data
		x, y := byte(nibbles[1]), byte(nibbles[2])
		n := byte(nibbles[3])
		return instructions.Display(x, y, n)
	}

	return nil
}
