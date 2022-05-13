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
				}
			}
		}
	case 0x1:
		address := uint16(nibbles[1])<<8 | uint16(nibbles[2])<<4 | uint16(nibbles[3])
		return fmt.Sprintf("Jump program counter to address %x", address)
	case 0x6:
		byteValue := byte(nibbles[2]<<4 | nibbles[3])
		registerIndex := byte(nibbles[1])
		return fmt.Sprintf("Set register v%x to %x", registerIndex, byteValue)
	case 0x7:
		byteValue := byte(nibbles[2]<<4 | nibbles[3])
		registerIndex := byte(nibbles[1])
		return fmt.Sprintf("Add %x to register v%x", byteValue, registerIndex)
	case 0xA:
		NNN := uint16(nibbles[1])<<8 | uint16(nibbles[2])<<4 | uint16(nibbles[3])
		return fmt.Sprintf("Set index register to %x", NNN)
	case 0xD: // DXYN - Draws a sprite at coordinate (X,Y) with N bytes of sprite data
		x, y := byte(nibbles[1]), byte(nibbles[2])
		n := byte(nibbles[3])
		return fmt.Sprintf("Draw sprite at (%d,%d) with %d bytes", x, y, n)
	}
	return "Unknown opcode"
}
