package indexregister

type IndexRegister uint16

func (r *IndexRegister) Add(value uint16) {
	*r += IndexRegister(value)
}

func (r *IndexRegister) Subtract(value uint16) {
	*r -= IndexRegister(value)
}

func (r *IndexRegister) Set(value uint16) {
	*r = IndexRegister(value)
}
