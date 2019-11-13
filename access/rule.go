package access

import "github.com/Workiva/go-datastructures/bitarray"

type Rule struct {
	bitarray.BitArray
}

func (r *Rule) Bits() bitarray.BitArray {
	return r.BitArray
}

func (r *Rule) Set(bit Bit) error {
	return r.SetBit(uint64(bit))
}

func NewRule() *Rule {
	return &Rule{
		BitArray: bitarray.NewSparseBitArray(),
	}
}
