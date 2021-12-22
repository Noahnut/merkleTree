package merkletree

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"log"
	"sync"
)

type MerkleTree struct {
	root       *node
	Leafs      sync.Map
	depth      int
	remainLeft int
}

type node struct {
	Left      *node
	Right     *node
	Parent    *node
	Isleaf    bool
	IsRoot    bool
	HashValue []byte
	Context   []byte
}

// Create the new Merkle Tree
func CreateMerkleTree() *MerkleTree {

	m := MerkleTree{
		depth: 0,
	}
	return &m
}

// Add new data block to the Merkle Tree
func (m *MerkleTree) AddNewBlock(context []byte) {
	var n node

	HashValue, err := m.calculateHash(context)

	if err != nil {
		log.Println("Hash write Failure: ", err.Error())
	}

	n.HashValue = HashValue

	n.Context = context
	n.Isleaf = true

	m.addNodeToTree(&n)

	m.Leafs.Store(hex.EncodeToString(n.HashValue), &n)
}

// Get the Tree root hash value
func (m *MerkleTree) GetRootHash() []byte {
	return m.root.HashValue
}

// Check Tree all the hash value is follow the Merkle Tree Properity
func (m *MerkleTree) CheckTreeCorrect() bool {
	return m.checkParentHashValue(m.root)
}

// Check the context exist and legal in the tree
// make sure all the hash from leaf to root is correct
func (m *MerkleTree) ContextValidator(context []byte) bool {

	ct, err := m.calculateHash(context)

	if err != nil {
		fmt.Println(err.Error())
		return false
	}

	Icn, exist := m.Leafs.Load(hex.EncodeToString(ct))

	if !exist {
		fmt.Println("Context not exist")
		return false
	}

	cn := Icn.(*node)

	cnParent := cn.Parent

	for cnParent != nil {

		cpH, err := m.getParentHashValue(cnParent)

		if err != nil {
			fmt.Println(err.Error())
			return false
		}

		if !bytes.Equal(cpH, cnParent.HashValue) {
			fmt.Println("Hash Function is not correct", hex.EncodeToString(cpH), hex.EncodeToString(cnParent.HashValue))
			return false
		}

		cnParent = cnParent.Parent
	}

	return true
}

// calculate Parent Hash Value from child
func (m *MerkleTree) calculateParentHashPath(n *node) {
	iter := n.Parent

	for {
		if iter == nil {
			return
		}

		HashValue, err := m.getParentHashValue(iter)

		if err != nil {
			log.Println("Hash write Failure: ", err.Error())
		}

		iter.HashValue = HashValue

		if iter.IsRoot {
			break
		}

		iter = iter.Parent
	}
}

// build the MerkleTree
func (m *MerkleTree) createNewTree(n *node) {
	rootNode := &node{
		Left:   n,
		Right:  nil,
		Parent: nil,
		Isleaf: false,
		IsRoot: true,
	}
	n.Isleaf, n.Parent = true, rootNode
	m.root = rootNode
	m.depth++
	m.remainLeft = (1 << m.depth) - 1
	m.calculateParentHashPath(n)
}

// add the node to the tree
// for the tree balance tree depth insrease only where left is full
func (m *MerkleTree) addNodeToTree(n *node) {
	if m.root == nil {
		m.createNewTree(n)
	} else {
		if m.remainLeft > 0 {
			currDep := 0
			iter := m.root
			rf := m.remainLeft
			for {
				if iter.Left == nil && iter.Right == nil && iter.Isleaf {
					np := node{}
					if iter.Parent.Left == iter {
						iter.Parent.Left = &np
					} else {
						iter.Parent.Right = &np
					}

					np.Parent = iter.Parent
					np.Left = iter
					np.Right = n
					n.Parent, iter.Parent = &np, &np
					m.remainLeft--
					break
				} else if iter.Right == nil {
					iter.Right = n
					n.Parent = iter
					break
				}

				if rf > (1<<(m.depth-currDep))/2 {
					iter = iter.Left
					currDep += 1
					rf -= (1 << (m.depth - currDep))
				} else {
					iter = iter.Right
					currDep += 1
				}
			}
		}
		m.remainLeft--
		if m.remainLeft <= 0 {
			m.depth++
			m.remainLeft = 1 << m.depth
		}
		m.calculateParentHashPath(n)
	}
}

func (m *MerkleTree) checkParentHashValue(p *node) bool {

	if p == nil || p.Isleaf {
		return true
	}

	cpH, err := m.getParentHashValue(p)

	if err != nil {
		return false
	}

	if !bytes.Equal(cpH, p.HashValue) {
		fmt.Println("Hash Function is not correct", hex.EncodeToString(cpH), hex.EncodeToString(p.HashValue))
		return false
	}

	m.checkParentHashValue(p.Left)
	m.checkParentHashValue(p.Right)
	return true
}

func (m *MerkleTree) getParentHashValue(p *node) ([]byte, error) {
	lh, rh := make([]byte, 0), make([]byte, 0)

	if p.Left != nil {
		lh = p.Left.HashValue
	}

	if p.Right != nil {
		rh = p.Right.HashValue
	}

	cpH, err := m.calculateHash(append(lh, rh...))

	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	return cpH, nil
}

func (m *MerkleTree) calculateHash(context []byte) ([]byte, error) {
	h := sha256.New()
	if _, err := h.Write(context); err != nil {
		return nil, err
	}

	return h.Sum(nil), nil
}

//helper function
func (m *MerkleTree) PrintCurrTree() {

	queue := make([]node, 0)
	queue = append(queue, *m.root)

	for len(queue) > 0 {
		q := queue[0]

		log.Println(q.IsRoot, q.Isleaf, hex.EncodeToString(q.HashValue), string(q.Context))

		if q.Parent != nil {
			log.Println(string(q.Context), hex.EncodeToString(q.Parent.HashValue))
		}

		if q.Left != nil {
			queue = append(queue, *q.Left)
		}

		if q.Right != nil {
			queue = append(queue, *q.Right)
		}

		queue = queue[1:]
	}
}

func (m *MerkleTree) checkLeafExist(hashString string) bool {
	_, ok := m.Leafs.Load(hashString)
	return ok
}

//helper function
func (m *MerkleTree) PrintlnAllLeafs() {
	m.Leafs.Range(func(key, value interface{}) bool {
		fmt.Println(key.(string), string(value.(*node).Context))
		return true
	})
}
