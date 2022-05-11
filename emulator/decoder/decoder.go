package decoder

import (
	"github.com/alperenbirol/chip-8/emulator/instructions"
	"github.com/alperenbirol/chip-8/emulator/memory"
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
	case 0x1: // 1nnn - JP addr
		// convert 3 nibbles to a 16 bit address
		address := uint16(nibbles[1])<<8 | uint16(nibbles[2])<<4 | uint16(nibbles[3])
		return instructions.Jump(memory.Address(address))
	case 0x6: // 6XNN: LD Vx, byte
		// convert 2 nibbles to a 8 bit address
		registerValue := byte(nibbles[2]<<4 | nibbles[3])
		register := byte(nibbles[1])
		// set VX(register) register to NN(registerValue)
		return instructions.SetRegister(register, registerValue)
	}

	return nil
}
