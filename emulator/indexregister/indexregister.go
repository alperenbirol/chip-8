package indexregister

import "github.com/alperenbirol/chip-8/emuconfig"

type IndexRegister emuconfig.Address

func (r *IndexRegister) Add(value uint16) {
	*r += IndexRegister(value)
}

func (r *IndexRegister) Subtract(value uint16) {
	*r -= IndexRegister(value)
}

func (r *IndexRegister) Set(value uint16) {
	*r = IndexRegister(value)
}

func (r *IndexRegister) Get() emuconfig.Address {
	return emuconfig.Address(*r)
}
