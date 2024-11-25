package bst

import (
	"sync"

	errors "github.com/imariom/nexusdb/pkg/errors"
	kv "github.com/imariom/nexusdb/pkg/kvpair"
)

// node represents a single node in the binary search tree that
// contains data as a key/value pair.
type node struct {
	// data is a pointer to where the key/value pair is stored.
	data *kv.KVPair

	// left is the pointer to the left node.
	left *node

	// right is the pointer to the right node.
	right *node
}

// BST represents the binary search tree.
type BST struct {
	// root is the root node of the tree.
	root *node

	// mu is the read and write mutex used to synchronize
	// ready and write operations in the BST tree.
	mu sync.RWMutex
}

// Insert inserts a new key/value pair in the tree or updates it if it exists
// and is not expired.
func (bst *BST) Insert(pair *kv.KVPair) error {
	bst.mu.Lock()
	defer bst.mu.Unlock()

	p, err := pair.Clone()
	if err != nil {
		return err
	} else if bst.root == nil {
		bst.root = &node{
			data: p,
		}
		return nil
	}
	return inserNode(bst.root, p)
}

// Get return a deep copy
func (bst *BST) Get(key []byte) (*kv.KVPair, error) {
	bst.mu.Lock()
	defer bst.mu.Unlock()
	return getNode(bst.root, key)
}

// InOrder traverses the tree in-order (left, root, right).
func (bst *BST) InOrder() []*kv.KVPair {
	// bst.mu.RLock()
	// defer bst.mu.RUnlock()

	// var result []*kv.KVPair
	// inOrderTraversal(bst.root, &result)

	// return result
	return nil
}

// Search searches for a non expired key/value pair in the BST tree.
func (bst *BST) Search(key []byte) bool {
	bst.mu.RLock()
	defer bst.mu.RUnlock()
	return searchNode(bst.root, key)
}

// Delete removes a key value pair from the tree if it exists.
func (bst *BST) Delete(key []byte) error {
	// bst.mu.Lock()
	// defer bst.mu.Unlock()
	// bst.root = deleteNode(bst.root, key)
	return nil
}

// inserNode inserts or updates a key/value pair in the tree.
// If the key/value pair exists and is expired it will return an error
func inserNode(n *node, p *kv.KVPair) error {
	pKey, _ := p.HashedKey()
	nKey, _ := n.data.HashedKey()

	if pKey < nKey {
		// insert a new node in the left of the current node
		if n.left == nil {
			n.left = &node{
				data: p,
			}
			return nil
		}
		return inserNode(n.left, p)
	} else if pKey > nKey {
		// insert a new node in the right of the current node
		if n.right == nil {
			n.right = &node{
				data: p,
			}
			return nil
		}
		return inserNode(n.right, p)
	}
	// ensure to update the current node if it already exists
	return updateNode(n, p)
}

// updateNode updates the node with the key/value pair passed.
func updateNode(n *node, p *kv.KVPair) error {
	if n == nil {
		return errors.ErrNodeIsNil
	}

	pkey, _ := p.HashedKey()
	nKey, _ := n.data.HashedKey()

	if pkey == nKey {
		return n.data.UpdateWith(p)
	} else if pkey < nKey {
		return updateNode(n.left, p)
	} else {
		return updateNode(n.right, p)
	}
}

// getNode returns the KVPair from the tree identified by the key k.
func getNode(n *node, k []byte) (*kv.KVPair, error) {
	if n == nil {
		return nil, errors.ErrKeyNotFound
	}

	key := kv.HashKey(k)
	nKey, _ := n.data.HashedKey()

	if key == nKey {
		pair, err := n.data.Clone()
		if err != nil {
			return nil, err
		}
		return pair, nil
	} else if key < nKey {
		return getNode(n.left, k)
	}
	return getNode(n.right, k)
}

// inOrderTraversal traverses the tree in-order (left, n, right).
func inOrderTraversal(n *node, r *[]*kv.KVPair) {
	// if n != nil {
	// 	inOrderTraversal(n.left, r)
	// 	if !n.data.Expired() {
	// 		*r = append(*r, n.data.Copy())
	// 	}
	// 	inOrderTraversal(n.right, r)
	// }
}

// searchNode traverses the BST tree trying to find given k.
func searchNode(n *node, k []byte) bool {
	if n == nil {
		return false
	}

	key := kv.HashKey(k)
	nKey, _ := n.data.HashedKey()

	if key == nKey {
		return true
	} else if key < nKey {
		return searchNode(n.left, k)
	} else {
		return searchNode(n.right, k)
	}
}

func deleteNode(n *node, k []byte) *node {
	// if n == nil {
	// 	return nil
	// }

	// pkey := kv.hashkey(k)
	// nkey := n.data.gethashedkey()

	// if pkey < nkey {
	// 	n.left = deletenode(n.left, k)
	// } else if pkey > nkey {
	// 	n.right = deletenode(n.right, k)
	// } else {
	// 	// node to be deleted is found
	// 	if n.left == nil {
	// 		return n.right
	// 	} else if n.right == nil {
	// 		return n.left
	// 	}
	// 	// node has two children
	// 	minright := findmin(n.right)
	// 	n.data = minright.data.copy()
	// 	k, _ := minright.data.getkey()
	// 	n.right = deletenode(n.right, k)
	// }

	// return n
	return nil
}

func findMin(n *node) *node {
	// current := n
	// for current.left != nil {
	// 	current = current.left
	// }
	// return current

	return nil
}
