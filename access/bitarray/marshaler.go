package bitarray

import (
	"encoding/json"

	"github.com/Workiva/go-datastructures/bitarray"
)

func NewMarshaler(ba bitarray.BitArray) Marshaler {
	return NewBitArrayHEXRaw(ba)
	//return NewBitArrayLZ4_HEX(ba)
	//return NewSliceUint64(ba)
	//return NewBitArrayB58Raw(ba)
}

// todo refactor
type Marshaler interface {
	bitarray.BitArray
	json.Marshaler
	json.Unmarshaler
}
