package programregister

type ProgramRegister byte

func (r *ProgramRegister) Add(value byte) {
	*r += ProgramRegister(value)
}

func (r *ProgramRegister) Subtract(value byte) {
	*r -= ProgramRegister(value)
}

func (r *ProgramRegister) Set(value byte) {
	*r = ProgramRegister(value)
}
