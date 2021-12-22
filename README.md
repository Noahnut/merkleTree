[![Go Reference](https://pkg.go.dev/badge/github.com/Noahnut/merkletree.svg)](https://pkg.go.dev/github.com/Noahnut/merkletree)

# MerkleTree Golang
MerkleTree is a hash tree which hash the node context from the leaf to the root. all the parent hash value has relation with its child node, this data struct can use for easily verify data set in the O(logN) time complexity.
In this implement have the add new leaf, get the root hash value, check the context is correct feature.


## Install 
```
go get github.com/Noahnut/merkletree
```


## Example Usage
```go
package main

import (
	"fmt"

	"github.com/Noahnut/merkletree"
)

func main() {
	m := merkletree.CreateMerkleTree()
	a := "testString"
	b := "testStringTwo"
	c := "testStringThree"
	d := "testStringFour"
	e := "testStringFive"

	// Add the new data to the tree
	m.AddNewBlock([]byte(a))
	m.AddNewBlock([]byte(b))
	m.AddNewBlock([]byte(c))
	m.AddNewBlock([]byte(d))
	m.AddNewBlock([]byte(e))

	// Verify the data is exist in the tree or not
	result := m.ContextValidator([]byte(e))

	// Get the root hash value 
	rootHash := m.GetRootHash()

	// Check tree from top to down hash value is correct
	result := m.CheckTreeCorrect()
}
```