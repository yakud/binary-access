package access

import (
	"fmt"
)

type Checker struct {
	tree *ChainTree
}

// HasAccess return true if rule have access to chain
func (c *Checker) HasAccess(chainName string, rule *Rule) bool {
	v, err := c.tree.GetVertexByName(chainName)
	if err != nil {
		fmt.Println(err)
		return false
	}

	return !v.Chain.And(rule.Bits()).IsEmpty()
}

func NewChecker(tree *ChainTree) *Checker {
	return &Checker{
		tree: tree,
	}
}
