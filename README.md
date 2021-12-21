# MerkleTree Golang
MerkleTree is a hash tree which hash from the leaf to the root, each hash value have the relation so can easy to verify the data from different source.
In this implement use the sha256 as the hash function and provide add new data block, verify data and get the root hash value feature.


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