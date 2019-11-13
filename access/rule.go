package access

import "github.com/Workiva/go-datastructures/bitarray"

type Rule struct {
	chainBits bitarray.BitArray
}

func (r *Rule) Bits() bitarray.BitArray {
	return r.chainBits
}

func (r *Rule) Set(bit Bit) error {
	return r.chainBits.SetBit(uint64(bit))
}

func NewRule() *Rule {
	return &Rule{
		chainBits: bitarray.NewSparseBitArray(),
	}
}
