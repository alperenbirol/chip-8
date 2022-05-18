package decoder

import (
	"fmt"

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
				case 0xE: // 00EE - Returns from a subroutine
					return instructions.ReturnFromSubroutine()
				}
			}
		}
	case 0x1: // 1NNN - Jumps program counter to address NNN
		// convert 3 nibbles to a 16 bit address
		address := uint16(nibbles[1])<<8 | uint16(nibbles[2])<<4 | uint16(nibbles[3])
		return instructions.Jump(emuconfig.Address(address))
	case 0x2: // 2NNN - Calls subroutine at NNN
		NNN := uint16(nibbles[1])<<8 | uint16(nibbles[2])<<4 | uint16(nibbles[3])
		return instructions.CallSubroutine(emuconfig.Address(NNN))
	case 0x3: // 3XNN - Skips the next instruction if register X equals NN
		byteValue := byte(nibbles[2]<<4 | nibbles[3])
		return instructions.SkipIfEqual(nibbles[1], byteValue)
	case 0x4: // 4XNN - Skips the next instruction if register X doesn't equal NN
		byteValue := byte(nibbles[2]<<4 | nibbles[3])
		return instructions.SkipIfNotEqual(nibbles[1], byteValue)
	case 0x5:
		vx, vy := nibbles[1], nibbles[2]
		return instructions.SkipIfRegistersHaveSameValue(vx, vy)
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
	case 0x9:
		vx, vy := nibbles[1], nibbles[2]
		return instructions.SkipIfRegistersHaveDifferentValue(vx, vy)
	case 0x8:
		switch nibbles[3] {
		case 0x0: // 8XY0 - Sets register X to register Y
			vx, vy := nibbles[1], nibbles[2]
			return instructions.SetRegisterToRegister(vx, vy)
		case 0x1: // 8XY1 - Sets register X to register X OR register Y
			vx, vy := nibbles[1], nibbles[2]
			return instructions.SetRegisterToRegisterOr(vx, vy)
		case 0x2: // 8XY2 - Sets register X to register X AND register Y
			vx, vy := nibbles[1], nibbles[2]
			return instructions.SetRegisterToRegisterAnd(vx, vy)
		case 0x3: // 8XY3 - Sets register X to register X XOR register Y
			vx, vy := nibbles[1], nibbles[2]
			return instructions.SetRegisterToRegisterXor(vx, vy)
		case 0x4: // 8XY4 - Adds register Y to register X. VF is set to 1 when there's a carry, and to 0 when there isn't
			vx, vy := nibbles[1], nibbles[2]
			return instructions.AddRegisterToRegister(vx, vy)
		case 0x5: // 8XY5 - VY is subtracted from VX. VF is set to 0 when there's a borrow, and 1 when there isn't
			vx, vy := nibbles[1], nibbles[2]
			return instructions.SubtractRegisterFromRegister(vx, vy)
		case 0x6: // 8XY6 - Shifts register X right by one. VF is set to the value of the least significant bit of VX before the shift
			vx, vy := nibbles[1], nibbles[2]
			return instructions.ShiftRegisterRight(vx, vy)
		case 0x7: // 8XY7 - Sets register X to VY minus VX. VF is set to 0 when there's a borrow, and 1 when there isn't
			vx, vy := nibbles[1], nibbles[2]
			return instructions.SubtractRegisterFromRegisterReverse(vx, vy)
		case 0xE: // 8XYE - Shifts register X left by one. VF is set to the value of the most significant bit of VX before the shift
			vx, vy := nibbles[1], nibbles[2]
			return instructions.ShiftRegisterLeft(vx, vy)
		}
	case 0xA: // ANNN - Sets index register to NNN
		// convert 3 nibbles to a 16 bit address
		NNN := uint16(nibbles[1])<<8 | uint16(nibbles[2])<<4 | uint16(nibbles[3])
		return instructions.SetIndexRegister(NNN)
	case 0xB: // BNNN - Jumps to address NNN plus V0
		NNN := uint16(nibbles[1])<<8 | uint16(nibbles[2])<<4 | uint16(nibbles[3])
		return instructions.JumpToAddressPlusRegister(NNN)
	case 0xC:
		NN := uint16(nibbles[2]<<4 | nibbles[3])
		return instructions.SetRegisterToRandom(nibbles[1], NN)
	case 0xD: // DXYN - Draws a sprite at coordinate (X,Y) with N bytes of sprite data
		x, y := byte(nibbles[1]), byte(nibbles[2])
		n := byte(nibbles[3])
		return instructions.Display(x, y, n)
	case 0xE:
		NN := byte(nibbles[2]<<4 | nibbles[3])
		if NN == 0x9E {
			return instructions.SkipIfKeyPressed(nibbles[1])
		}
		if NN == 0xA1 {
			return instructions.SkipIfKeyNotPressed(nibbles[1])
		}
	case 0xF:
		NN := uint16(nibbles[2])<<4 | uint16(nibbles[3])
		switch NN {
		case 0x07:
			return instructions.SetRegisterToDelayTimer(nibbles[1])
		case 0x0A:
			return instructions.WaitForKeyPress(nibbles[1])
		case 0x15:
			return instructions.SetDelayTimer(nibbles[1])
		case 0x18:
			return instructions.SetSoundTimer(nibbles[1])
		case 0x1E:
			return instructions.AddToIndexRegister(nibbles[1])
		case 0x29:
			return instructions.SetIndexRegisterToFont(nibbles[1])
		case 0x33:
			return instructions.StoreBCD(nibbles[1])
		case 0x55:
			return instructions.StoreRegistersAtMemory(nibbles[1])
		case 0x65:
			return instructions.FillRegistersFromMemory(nibbles[1])
		}
	}
	fmt.Printf("0x%x%x%x%x", nibbles[0], nibbles[1], nibbles[2], nibbles[3])
	return nil
}
