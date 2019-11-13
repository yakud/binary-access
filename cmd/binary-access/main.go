package main

import (
	"fmt"
	"log"
	"time"

	"github.com/yakud/binary-access/access"
)

func main() {
	// generate tree
	tree, err := generateTree()
	if err != nil {
		log.Fatal(err)
	}

	accessChecker := access.NewChecker(tree)

	// 2 users
	userReadAccess := []*access.Rule{
		access.NewRule(), // user 0 read rules
		access.NewRule(), // user 1 read rules
	}

	// set user 0 read access to "a.c" subtree
	if err := tree.SetRuleBitByChainName(userReadAccess[0], "a.c"); err != nil {
		log.Fatal(err)
	}
	if err := tree.SetRuleBitByChainName(userReadAccess[0], "a.b.g"); err != nil {
		log.Fatal(err)
	}

	// set user 1 read access to "a.c.d.e" subtree
	if err := tree.SetRuleBitByChainName(userReadAccess[1], "a.c.d.e"); err != nil {
		log.Fatal(err)
	}

	runChecks(accessChecker, userReadAccess)

	since := time.Now()
	treeJson, err := tree.MarshalJSON()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("tree marshaled to JSON for:", time.Since(since))
	fmt.Println(string(treeJson))

	unmarshaledTree := access.NewTree(0)
	since = time.Now()
	if err := unmarshaledTree.UnmarshalJSON(treeJson); err != nil {
		log.Fatal(err)
	}
	fmt.Println("tree unmarshaled to struct for:", time.Since(since))
	accessChecker = access.NewChecker(unmarshaledTree)

	fmt.Println("RUN unmarhaled checks")
	runChecks(accessChecker, userReadAccess)
}

func runChecks(accessChecker *access.Checker, userReadAccess []*access.Rule) {
	// Check user 0
	if !accessChecker.HasAccess("a.c.d", userReadAccess[0]) {
		log.Fatal("FAILED! User 0 should have access to a.c.d")
	} else {
		fmt.Println("check done!")
	}

	if !accessChecker.HasAccess("a.c.d.f", userReadAccess[0]) {
		log.Fatal("FAILED! User 0 should have access to a.c.d.f")
	} else {
		fmt.Println("check done!")
	}

	if accessChecker.HasAccess("a", userReadAccess[0]) {
		log.Fatal("FAILED! User 0 should not have access to a")
	} else {
		fmt.Println("check done!")
	}

	if accessChecker.HasAccess("a.b", userReadAccess[0]) {
		log.Fatal("FAILED! User 0 should not have access to a.b")
	} else {
		fmt.Println("check done!")
	}

	if !accessChecker.HasAccess("a.b.g", userReadAccess[0]) {
		log.Fatal("FAILED! User 0 should have access to a.b.g")
	} else {
		fmt.Println("check done!")
	}

	// Check user 1
	if accessChecker.HasAccess("a", userReadAccess[1]) {
		log.Fatal("FAILED! User 1 should not have access to a")
	} else {
		fmt.Println("check done!")
	}
	if accessChecker.HasAccess("a.c.d", userReadAccess[1]) {
		log.Fatal("FAILED! User 1 should not have access to a.c.d")
	} else {
		fmt.Println("check done!")
	}
	if accessChecker.HasAccess("a.b.g", userReadAccess[1]) {
		log.Fatal("FAILED! User 1 should not have access to a.b.g")
	} else {
		fmt.Println("check done!")
	}
	if !accessChecker.HasAccess("a.c.d.e", userReadAccess[1]) {
		log.Fatal("FAILED! User 1 should have access to a.c.d.e")
	} else {
		fmt.Println("check done!")
	}
}

//
// Generate tree like:
//                e
//               /
//          c - d - f
//         /
// (root) a - b - g
//
func generateTree() (*access.ChainTree, error) {
	tree := access.NewTree(0)
	a, err := tree.Add("a", nil)
	if err != nil {
		return nil, err
	}

	b, err := tree.Add("b", a)
	if err != nil {
		return nil, err
	}

	c, err := tree.Add("c", a)
	if err != nil {
		return nil, err
	}

	d, err := tree.Add("d", c)
	if err != nil {
		return nil, err
	}

	_, err = tree.Add("e", d)
	if err != nil {
		return nil, err
	}

	_, err = tree.Add("f", d)
	if err != nil {
		return nil, err
	}

	_, err = tree.Add("g", b)
	if err != nil {
		return nil, err
	}

	return tree, nil
}
