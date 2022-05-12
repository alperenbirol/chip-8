package emuconfig

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

type Address uint16

type Pixel bool

type Pixels [DISPLAY_HEIGHT][DISPLAY_WIDTH]Pixel
