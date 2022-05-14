package decoder

import (
	"fmt"

	"github.com/alperenbirol/chip-8/emuconfig"
)

func Description(opcode *emuconfig.Opcode) string {
	nibbles := opcode.GetNibbles()
	switch nibbles[0] {
	case 0x0:
		switch nibbles[1] {
		case 0x0:
			switch nibbles[2] {
			case 0xE:
				switch nibbles[3] {
				case 0x0:
					return "Clear screen"
				case 0xE:
					return "Return from subroutine"
				}
			}
		}
	case 0x1:
		address := uint16(nibbles[1])<<8 | uint16(nibbles[2])<<4 | uint16(nibbles[3])
		return fmt.Sprintf("Jump program counter to address 0x%x", address)
	case 0x2: // 2NNN - Calls subroutine at NNN
		NNN := uint16(nibbles[1])<<8 | uint16(nibbles[2])<<4 | uint16(nibbles[3])
		return fmt.Sprintf("Call subroutine at address 0x%x", NNN)
	case 0x3:
		byteValue := byte(nibbles[2]<<4 | nibbles[3])
		return fmt.Sprintf("Skip next instruction if register v%x's value equals 0x%x", nibbles[1], byteValue)
	case 0x4:
		byteValue := byte(nibbles[2]<<4 | nibbles[3])
		return fmt.Sprintf("Skip next instruction if register v%x's value != 0x%x", nibbles[1], byteValue)
	case 0x5:
		vx, vy := nibbles[1], nibbles[2]
		return fmt.Sprintf("Skip next instruction if register v%x's value == register v%x's value", vx, vy)
	case 0x6:
		byteValue := byte(nibbles[2]<<4 | nibbles[3])
		registerIndex := byte(nibbles[1])
		return fmt.Sprintf("Set register v%x to 0x%x", registerIndex, byteValue)
	case 0x7:
		byteValue := byte(nibbles[2]<<4 | nibbles[3])
		registerIndex := byte(nibbles[1])
		return fmt.Sprintf("Add 0x%x to register v%x", byteValue, registerIndex)
	case 0x8:
		switch nibbles[3] {
		case 0x0:
			vx, vy := nibbles[1], nibbles[2]
			return fmt.Sprintf("Set register v%x to register v%x", vx, vy)
		case 0x1:
			vx, vy := nibbles[1], nibbles[2]
			return fmt.Sprintf("Set register v%x to register v%x | register v%x", vx, vy, vx)
		case 0x2:
			vx, vy := nibbles[1], nibbles[2]
			return fmt.Sprintf("Set register v%x to register v%x & register v%x", vx, vy, vx)
		case 0x3:
			vx, vy := nibbles[1], nibbles[2]
			return fmt.Sprintf("Set register v%x to register v%x ^ register v%x", vx, vy, vx)
		case 0x4:
			vx, vy := nibbles[1], nibbles[2]
			return fmt.Sprintf("Add register v%x to register v%x, store in register v%x", vy, vx, vx)
		case 0x5: // 8XY5 - Set register vx to register vx - register vy
			vx, vy := nibbles[1], nibbles[2]
			return fmt.Sprintf("Set register v%x to register v%x - register v%x", vx, vx, vy)
		case 0x6:
			vx, vy := nibbles[1], nibbles[2]
			return fmt.Sprintf("Set register v%x to register v%x >> 1, store in register v%x", vy, vx, vx)
		case 0x7:
			vx, vy := nibbles[1], nibbles[2]
			return fmt.Sprintf("Set register v%x to register v%x - register v%x", vx, vy, vx)
		case 0xE:
			vx, vy := nibbles[1], nibbles[2]
			return fmt.Sprintf("Set register v%x to register v%x << 1, store in register v%x", vy, vx, vx)
		}
	case 0x9:
		vx, vy := nibbles[1], nibbles[2]
		return fmt.Sprintf("Skip next instruction if register v%x's value != register v%x's value", vx, vy)
	case 0xA:
		NNN := uint16(nibbles[1])<<8 | uint16(nibbles[2])<<4 | uint16(nibbles[3])
		return fmt.Sprintf("Set index register to 0x%x", NNN)
	case 0xB:
		NNN := uint16(nibbles[1])<<8 | uint16(nibbles[2])<<4 | uint16(nibbles[3])
		return fmt.Sprintf("Jump program counter to address 0x%x + register v0", NNN)
	case 0xC:
		NN := uint16(nibbles[2]<<4 | nibbles[3])
		return fmt.Sprintf("Set register v%x to random byte & 0x%x", nibbles[1], NN)
	case 0xD: // DXYN - Draws a sprite at coordinate (X,Y) with N bytes of sprite data
		x, y := byte(nibbles[1]), byte(nibbles[2])
		n := byte(nibbles[3])
		return fmt.Sprintf("Draw sprite at (%d,%d) with %d bytes", x, y, n)
	case 0xE:
		NN := byte(nibbles[2]<<4 | nibbles[3])
		if NN == 0xA1 {
			return fmt.Sprintf("Skip next instruction if key 0x%x is not pressed", nibbles[1])
		}
	case 0xF:
		NN := uint16(nibbles[2])<<4 | uint16(nibbles[3])
		switch NN {
		case 0x1E:
			return fmt.Sprintf("Add register v%x to index register", nibbles[1])
		case 0x07:
			return "Set register vx to delay timer value"
		case 0x15:
			return fmt.Sprintf("Set delay timer to value of register v%x", nibbles[1])
		case 0x18:
			return fmt.Sprintf("Set sound timer to value of register v%x", nibbles[1])
		case 0x29:
			return fmt.Sprintf("Set index register to address of font in v%x", nibbles[1])
		case 0x33:
			return fmt.Sprintf("Store the binary-coded decimal representation of v%x at the addresses I, I+1, and I+2", nibbles[1])
		case 0x55:
			return fmt.Sprintf("Store registers v0 to v%x in memory starting at index register", nibbles[1])
		case 0x65:
			return fmt.Sprintf("Read registers v0 to v%x from memory starting at index register", nibbles[1])
		}
	}
	fmt.Printf("0x%x%x%x%x", nibbles[0], nibbles[1], nibbles[2], nibbles[3])
	return "Unknown opcode"
}
