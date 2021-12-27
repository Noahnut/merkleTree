package merkletree

import (
	"strconv"
	"testing"
)

func BenchmarkSimpleTree(b *testing.B) {
	mt := CreateMerkleTree()

	for i := 0; i < b.N; i++ {
		mt.AddNewBlock([]byte(strconv.Itoa(i)))
	}

	for i := 0; i < b.N; i++ {
		result := mt.ContextValidator([]byte(strconv.Itoa(i)))

		if !result {
			b.Error("Hash Fail")
			return
		}
	}
}

func BenchmarkCompareDifferTree(b *testing.B) {
	first := CreateMerkleTree()
	Second := CreateMerkleTree()

	for i := 0; i < b.N; i++ {
		first.AddNewBlock([]byte(strconv.Itoa(i)))
	}

	for i := b.N; i >= 0; i-- {
		Second.AddNewBlock([]byte(strconv.Itoa(i)))
	}

	first.GetDifferentContextFromTree(Second)
}
