package access

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	"github.com/Workiva/go-datastructures/bitarray"
)

type TreeID uint32

type ChainTree struct {
	id      TreeID
	lastBit Bit
	chains  []*Chain
	names   map[string]Bit
}

type marshaledTree struct {
	Id      TreeID         `json:"id"`
	LastBit Bit            `json:"last_bit"`
	Chains  []*Chain       `json:"chains"`
	Names   map[string]Bit `json:"names"`
}

func (t *ChainTree) Add(name string, parent *Chain) (*Chain, error) {
	v := &Chain{
		Bit: t.lastBit,
	}

	var parentName string
	if parent == nil {
		if len(t.chains) != 0 {
			return nil, errors.New("head already allocated and vertex parent is empty. you can make new tree")
		}
		v.Chain = bitarray.NewSparseBitArray()
		//v.Chain = bitarray.NewBitArray(1000000)
		v.Name = name
	} else {
		v.Chain = bitarray.NewSparseBitArray().Or(parent.Chain)
		//v.Chain = bitarray.NewBitArray(1000000).Or(parent.Chain)
		v.Name = t.GenerateName(parent.Name, name)
	}

	if err := v.Chain.SetBit(uint64(v.Bit)); err != nil {
		return nil, err
	}

	vertexName := t.GenerateName(parentName, name)
	if _, err := t.GetVertexByName(vertexName); err == nil {
		return nil, fmt.Errorf("vertex with name %s already allocated", vertexName)
	}

	t.names[v.Name] = v.Bit
	t.chains = append(t.chains, v)

	t.lastBit++

	return v, nil
}

func (t *ChainTree) GenerateName(parts ...string) string {
	return strings.Join(parts, ".")
}

func (t *ChainTree) GetVertexBitByName(name string) (Bit, error) {
	bit, ok := t.names[name]
	if !ok {
		return 0, fmt.Errorf("vertex %s not found", name)
	}

	return bit, nil
}

func (t *ChainTree) GetVertexByName(name string) (*Chain, error) {
	bit, ok := t.names[name]
	if !ok {
		return nil, fmt.Errorf("vertex %s not found", name)
	}

	return t.GetVertexByBit(bit)
}

func (t *ChainTree) GetVertexByBit(bit Bit) (*Chain, error) {
	if int(bit) >= len(t.chains) {
		return nil, fmt.Errorf("vertex bit %d not found", bit)
	}

	v := t.chains[bit]
	if v == nil {
		return nil, fmt.Errorf("vertex bit %d not found", bit)
	}

	return v, nil
}

func (t *ChainTree) SetRuleBitByChainName(rule *Rule, name string) error {
	bit, err := t.GetVertexBitByName(name)
	if err != nil {
		return err
	}

	return rule.Set(bit)
}

func (t *ChainTree) MarshalJSON() ([]byte, error) {
	return json.Marshal(&marshaledTree{
		Id:      t.id,
		LastBit: t.lastBit,
		Chains:  t.chains,
		Names:   t.names,
	})
}

func (t *ChainTree) UnmarshalJSON(b []byte) error {
	marshaledTree := marshaledTree{
		Chains: t.chains,
		Names:  t.names,
	}

	if err := json.Unmarshal(b, &marshaledTree); err != nil {
		return err
	}

	t.id = marshaledTree.Id
	t.lastBit = marshaledTree.LastBit
	t.chains = marshaledTree.Chains
	t.names = marshaledTree.Names

	return nil
}

func NewTree(id TreeID) *ChainTree {
	return &ChainTree{
		id:     id,
		chains: make([]*Chain, 0), // todo size
		names:  make(map[string]Bit),
	}
}
