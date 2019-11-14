package bitarray

import (
	"github.com/gramework/common/b58"

	"github.com/Workiva/go-datastructures/bitarray"
	_ "github.com/gramework/common/b58"
)

type bitArrayB58Raw struct {
	bitarray.BitArray
}

func (c *bitArrayB58Raw) MarshalJSON() ([]byte, error) {
	if chainBin, err := bitarray.Marshal(c.BitArray); err != nil {
		return nil, err
	} else {
		return []byte(`"` + b58.Encode(chainBin) + `"`), nil // todo efficient hex encoding
	}
}

func (c *bitArrayB58Raw) UnmarshalJSON(data []byte) error {
	chainBin, err := b58.Decode(string(data[1 : len(data)-1])) // todo efficient hex decoding
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

func NewBitArrayB58Raw(ba bitarray.BitArray) Marshaler {
	return &bitArrayB58Raw{BitArray: ba}
}
