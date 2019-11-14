package bitarray

import (
	"encoding/json"

	"github.com/Workiva/go-datastructures/bitarray"
)

type bitArraySliceUint64 struct {
	bitarray.BitArray
}

func (c *bitArraySliceUint64) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.ToNums())
}

func (c *bitArraySliceUint64) UnmarshalJSON(data []byte) error {
	b := make([]uint64, 0)
	if err := json.Unmarshal(data, &b); err != nil {
		return err
	}

	c.BitArray = bitarray.NewSparseBitArray()
	for _, v := range b {
		if err := c.BitArray.SetBit(v); err != nil {
			return err
		}
	}

	return nil
}

func NewSliceUint64(ba bitarray.BitArray) Marshaler {
	return &bitArraySliceUint64{BitArray: ba}
}
