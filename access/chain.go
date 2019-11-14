package access

import (
	"encoding/json"

	"github.com/Workiva/go-datastructures/bitarray"

	_ "github.com/klauspost/compress/zstd"
	bitarrayMarshaler "github.com/yakud/binary-access/access/bitarray"
)

type Bit uint64

type Chain struct {
	TreeID TreeID
	Bit    Bit
	Name   string
	Chain  bitarray.BitArray
}

// todo: serialize original Chain
type marshaledChain struct {
	TreeID TreeID                      `json:"tree_id"`
	Bit    Bit                         `json:"bit"`
	Name   string                      `json:"name"`
	Chain  bitarrayMarshaler.Marshaler `json:"chain"`
}

func (c *Chain) MarshalJSON() ([]byte, error) {
	marshaledChain := marshaledChain{
		TreeID: c.TreeID,
		Bit:    c.Bit,
		Name:   c.Name,
		Chain:  bitarrayMarshaler.NewMarshaler(c.Chain),
	}

	return json.Marshal(&marshaledChain)
}

func (c *Chain) UnmarshalJSON(b []byte) error {
	marshaledChain := marshaledChain{
		Chain: bitarrayMarshaler.NewMarshaler(bitarray.NewSparseBitArray()),
	}
	if err := json.Unmarshal(b, &marshaledChain); err != nil {
		return err
	}

	c.TreeID = marshaledChain.TreeID
	c.Bit = marshaledChain.Bit
	c.Name = marshaledChain.Name
	c.Chain = marshaledChain.Chain

	return nil
}
