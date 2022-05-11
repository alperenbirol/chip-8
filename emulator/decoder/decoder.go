package decoder

import (
	"github.com/alperenbirol/chip-8/emulator/decoder/instructions"
)

type Opcode [2]byte

func (code *Opcode) Opcode() uint16 {
	return uint16(code[0])<<8 | uint16(code[1])
}

type Nibble byte

func (code *Opcode) GetNibbles() [4]Nibble {
	return [4]Nibble{
		Nibble(code[0] >> 4),
		Nibble(code[0] & 0x0F),
		Nibble(code[1] >> 4),
		Nibble(code[1] & 0x0F),
	}
}

func (code *Opcode) Decode() instructions.Instruction {
	nibbles := code.GetNibbles()
	switch nibbles[0] {
	case 0x0:
		switch nibbles[1] {
		case 0x0:
			switch nibbles[2] {
			case 0xE:
				switch nibbles[3] {
				case 0x0:
					return instructions.ClearScreen()
				}
			}
		}
	}

	return nil
}
