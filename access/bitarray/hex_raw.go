package bitarray

import (
	"encoding/hex"

	"github.com/Workiva/go-datastructures/bitarray"
)

type bitArrayHEXRaw struct {
	bitarray.BitArray
}

func (c *bitArrayHEXRaw) MarshalJSON() ([]byte, error) {
	if chainBin, err := bitarray.Marshal(c.BitArray); err != nil {
		return nil, err
	} else {
		return []byte(`"` + hex.EncodeToString(chainBin) + `"`), nil // todo efficient hex encoding
	}
}

func (c *bitArrayHEXRaw) UnmarshalJSON(data []byte) error {
	chainBin, err := hex.DecodeString(string(data[1 : len(data)-1])) // todo efficient hex decoding
	if err != nil {
		return err
	}

	chain, err := bitarray.Unmarshal(chainBin)
	if err != nil {
		return err
	}
	c.BitArray = chain

	return nil
}

func NewBitArrayHEXRaw(ba bitarray.BitArray) Marshaler {
	return &bitArrayHEXRaw{BitArray: ba}
}
