package merkletree

import (
	"testing"
)

func TestCreateNewTree(t *testing.T) {
	mt := CreateMerkleTree()

	a := "testString"
	b := "testStringTwo"
	c := "testStringThree"
	d := "testStringFour"
	e := "testStringFive"
	f := "testStringSix"
	h := "testStringSeven"
	i := "testStringEight"

	mt.AddNewBlock([]byte(a))
	mt.AddNewBlock([]byte(b))

	mt.AddNewBlock([]byte(c))

	mt.AddNewBlock([]byte(d))
	mt.AddNewBlock([]byte(e))
	mt.AddNewBlock([]byte(f))
	mt.AddNewBlock([]byte(h))
	mt.AddNewBlock([]byte(i))

	result := mt.ContextValidator([]byte(i))

	if !result {
		t.Error("Validator Fail")
	}

	result = mt.ContextValidator([]byte(a))

	if !result {
		t.Error("Validator Fail")
	}

}
