package merkletree

import (
	"crypto/sha256"
	"encoding/hex"
	"strconv"
	"testing"
)

func HashCalculator(context []byte) ([]byte, error) {
	h := sha256.New()
	if _, err := h.Write(context); err != nil {
		return nil, err
	}

	return h.Sum(nil), nil
}

func TestCreateNewTree(t *testing.T) {
	mt := CreateMerkleTree()

	testStringOne := "testString"
	testStringOneExpectSHA256 := "4acf0b39d9c4766709a3689f553ac01ab550545ffa4544dfc0b2cea82fba02a3"
	testStringTwo := "testStringTwo"
	StringTwoExpectSHA256 := "a0f5924b97a3957686098083f696f24917955b844813f98563bb432f74869c25"
	testStringThree := "testStringThree"
	testStringThreeExpectSHA256 := "0fec6c35e72054934c9ed3c65cabcd1ac0fe004ded1e84a0548eef3d6a4fc53f"
	testStringFour := "testStringFour"
	testStringFourExpectSHA256 := "4cb932f7d45b7eb5de3419c9d2efacacc96efe684c504975681e19353e94e496"

	mt.AddNewBlock([]byte(testStringOne))
	exist := mt.checkLeafExist(testStringOneExpectSHA256)
	if !exist {
		t.Error("SHA256 Fail")

	}

	mt.AddNewBlock([]byte(testStringTwo))
	exist = mt.checkLeafExist(StringTwoExpectSHA256)
	if !exist {
		t.Error("SHA256 Fail")

	}

	mt.AddNewBlock([]byte(testStringThree))
	exist = mt.checkLeafExist(testStringThreeExpectSHA256)
	if !exist {
		t.Error("SHA256 Fail")

	}

	mt.AddNewBlock([]byte(testStringFour))
	exist = mt.checkLeafExist(testStringFourExpectSHA256)
	if !exist {
		t.Error("SHA256 Fail")

	}
	result := mt.ContextValidator([]byte(testStringOne))

	if !result {
		t.Error("Validator Fail")
	}

	result = mt.ContextValidator([]byte(testStringTwo))

	if !result {
		t.Error("Validator Fail")
	}

	result = mt.ContextValidator([]byte(testStringThree))

	if !result {
		t.Error("Validator Fail")
	}

	result = mt.ContextValidator([]byte(testStringFour))

	if !result {
		t.Error("Validator Fail")
	}

	result = mt.CheckTreeCorrect()

	if !result {
		t.Error("Validator Fail")
	}
}

func TestOddLeaf(t *testing.T) {
	mt := CreateMerkleTree()

	testStringOne := "testString"
	testStringOneExpectSHA256 := "4acf0b39d9c4766709a3689f553ac01ab550545ffa4544dfc0b2cea82fba02a3"
	testStringTwo := "testStringTwo"
	StringTwoExpectSHA256 := "a0f5924b97a3957686098083f696f24917955b844813f98563bb432f74869c25"
	testStringThree := "testStringThree"
	testStringThreeExpectSHA256 := "0fec6c35e72054934c9ed3c65cabcd1ac0fe004ded1e84a0548eef3d6a4fc53f"

	mt.AddNewBlock([]byte(testStringOne))
	exist := mt.checkLeafExist(testStringOneExpectSHA256)
	if !exist {
		t.Error("SHA256 Fail")

	}

	mt.AddNewBlock([]byte(testStringTwo))
	exist = mt.checkLeafExist(StringTwoExpectSHA256)
	if !exist {
		t.Error("SHA256 Fail")

	}

	mt.AddNewBlock([]byte(testStringThree))
	exist = mt.checkLeafExist(testStringThreeExpectSHA256)
	if !exist {
		t.Error("SHA256 Fail")
	}

	result := mt.ContextValidator([]byte(testStringOne))

	if !result {
		t.Error("Validator Fail")
	}

	result = mt.ContextValidator([]byte(testStringTwo))

	if !result {
		t.Error("Validator Fail")
	}

	result = mt.ContextValidator([]byte(testStringThree))

	if !result {
		t.Error("Validator Fail")
	}

	result = mt.CheckTreeCorrect()

	if !result {
		t.Error("Validator Fail")
	}
}

func TestManyLeaf(t *testing.T) {
	mt := CreateMerkleTree()

	type testPair struct {
		Leaf string
		hash []byte
	}

	testLeafs := make([]testPair, 0)

	for i := 0; i < 100000; i++ {
		tp := testPair{
			Leaf: "testString@" + strconv.Itoa(i),
		}

		hash, err := HashCalculator([]byte(tp.Leaf))

		if err != nil {
			t.Error("Hash Fail")
			return
		}

		tp.hash = hash

		testLeafs = append(testLeafs, tp)
	}

	for i := 0; i < len(testLeafs); i++ {
		mt.AddNewBlock([]byte(testLeafs[i].Leaf))
	}

	for i := 0; i < len(testLeafs); i++ {
		exist := mt.checkLeafExist(hex.EncodeToString(testLeafs[i].hash))

		if !exist {
			t.Error("SHA256 Fail")
		}
	}

	for i := 0; i < len(testLeafs); i++ {
		result := mt.ContextValidator([]byte(testLeafs[i].Leaf))

		if !result {
			t.Error("Validator Fail")
		}
	}

	result := mt.CheckTreeCorrect()

	if !result {
		t.Error("Validator Fail")
	}
}

func TestFindTreeDifferentContext_OneDifferent(t *testing.T) {
	firstTree := CreateMerkleTree()

	testStringOne := "testString"
	testStringTwo := "testStringTwo"
	testStringThree := "testStringThree"
	testStringFour := "testStringFour"
	testStringFive := "testStringFive"
	testStringSix := "testStringSix"

	firstTree.AddNewBlock([]byte(testStringOne))
	firstTree.AddNewBlock([]byte(testStringTwo))
	firstTree.AddNewBlock([]byte(testStringThree))
	firstTree.AddNewBlock([]byte(testStringFour))
	firstTree.AddNewBlock([]byte(testStringFive))
	firstTree.AddNewBlock([]byte(testStringSix))

	SecondTree := CreateMerkleTree()

	testStringOne = "testString"
	testStringTwo = "testStringTwo"
	testStringThree = "testStringThree"
	testStringFour = "testStringFour"
	testStringFive = "testStringFive"
	testStringSix = "testStringSix_diff"

	SecondTree.AddNewBlock([]byte(testStringOne))
	SecondTree.AddNewBlock([]byte(testStringTwo))
	SecondTree.AddNewBlock([]byte(testStringThree))
	SecondTree.AddNewBlock([]byte(testStringFour))
	SecondTree.AddNewBlock([]byte(testStringFive))
	SecondTree.AddNewBlock([]byte(testStringSix))

	diffContext := firstTree.GetDifferentContextFromTree(SecondTree)

	if len(diffContext) != 1 {
		t.Error("Wrong different context length")
	}

	if string(diffContext[0]) != testStringSix {
		t.Error("Wrong different context")
	}
}
