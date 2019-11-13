package access

import (
	"encoding/json"

	"github.com/Workiva/go-datastructures/bitarray"
	_ "github.com/klauspost/compress/zstd"
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
	TreeID   TreeID `json:"tree_id"`
	Bit      Bit    `json:"bit"`
	Name     string `json:"name"`
	ChainHEX string `json:"chain_hex"`
}

func (c *Chain) MarshalJSON() ([]byte, error) {
	marshaledChain := marshaledChain{
		TreeID: c.TreeID,
		Bit:    c.Bit,
		Name:   c.Name,
	}

	var err error
	chainHexBytes, err := BitArrayEncode(c.Chain)
	if err != nil {
		return nil, err
	}
	marshaledChain.ChainHEX = string(chainHexBytes)

	return json.Marshal(&marshaledChain)
}

func (c *Chain) UnmarshalJSON(b []byte) error {
	marshaledChain := marshaledChain{}
	if err := json.Unmarshal(b, &marshaledChain); err != nil {
		return err
	}

	var err error
	c.Chain, err = BitArrayDecode([]byte(marshaledChain.ChainHEX))
	if err != nil {
		return err
	}

	c.TreeID = marshaledChain.TreeID
	c.Bit = marshaledChain.Bit
	c.Name = marshaledChain.Name

	return nil
}
