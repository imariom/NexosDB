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
//
// Parameters:
//   - pair (KVPair): The key/value pair to be inserted/updated in the tree.
//     After sucessfull execution of the Insert method the argument should not be used further.
func (bst *BST) Insert(pair kv.KVPair) error {
	var err error

	bst.mu.Lock()
	defer bst.mu.Unlock()

	if bst.root == nil {
		bst.root = &node{
			data: pair.Copy(),
		}
	} else {
		err = inserNode(bst.root, pair)
	}

	return err
}

// Search searches for a non expired key/value pair in the BST tree.
func (bst *BST) Search(key []byte) bool {
	bst.mu.RLock()
	defer bst.mu.RUnlock()
	return searchNode(bst.root, key)
}

// Get return a deep copy
func (bst *BST) Get(key []byte) (*kv.KVPair, error) {
	bst.mu.Lock()
	defer bst.mu.Unlock()
	return getNode(bst.root, key)
}

// InOrder traverses the tree in-order (left, root, right).
func (bst *BST) InOrder() []*kv.KVPair {
	bst.mu.RLock()
	defer bst.mu.RUnlock()

	var result []*kv.KVPair
	inOrderTraversal(bst.root, &result)

	return result
}

// Deletion
// TODO: if key/value pair exists in the tree and expired remove it from the tree
func (bst *BST) Delete(key []byte) {
	bst.mu.Lock()
	defer bst.mu.Unlock()
	bst.root = deleteNode(bst.root, key)
}

// inserNode inserts or updates a key/value pair in the tree.
// If the key/value pair exists and is expired it will return an error
func inserNode(n *node, p kv.KVPair) error {
	if p.Expired() {
		return errors.ErrKeyExpired
	}

	pKey := p.GetHashedKey()
	nKey := n.data.GetHashedKey()
	var err error

	if pKey < nKey {
		// insert a new node in the left of the current node
		if n.left == nil {
			n.left = &node{
				data: p.Copy(),
			}
		} else {
			err = inserNode(n.left, p)
		}
	} else if pKey > nKey {
		// insert a new node in the right of the current node
		if n.right == nil {
			n.right = &node{
				data: p.Copy(),
			}
		} else {
			err = inserNode(n.right, p)
		}
	} else {
		// ensure to update the current node if it already exists
		err = updateNode(n, p)
	}

	return err
}

// updateNode updates the node with the key/value pair passed.
//
// Paramters:
//   - n (*node): The current node in the BST tree.
//   - p (KVPair): The key/value pair to update the node.
func updateNode(n *node, p kv.KVPair) error {
	if n == nil {
		return errors.ErrNodeIsNil
	}

	if p.Expired() {
		return errors.ErrKeyExpired
	}

	pkey := p.GetHashedKey()
	nKey := n.data.GetHashedKey()
	var err error

	if pkey == nKey {
		n.data = p.Copy()
	} else if pkey < nKey {
		err = updateNode(n.left, p)
	} else {
		err = updateNode(n.right, p)
	}

	return err
}

// searchNode traverses the BST tree trying to find given k.
//
// Parameters:
//   - n (*node): The current node in the BST tree.
//   - k ([]byte): The key to search in the BST tree.
func searchNode(n *node, k []byte) bool {
	if n == nil {
		return false
	}

	key := kv.HashKey(k)
	nKey := n.data.GetHashedKey()

	if key == nKey {
		return !n.data.Expired()
	} else if key < nKey {
		return searchNode(n.left, k)
	} else {
		return searchNode(n.right, k)
	}
}

func getNode(n *node, k []byte) (*kv.KVPair, error) {
	if n == nil {
		return nil, errors.ErrNodeIsNil
	}

	key := kv.HashKey(k)
	nKey := n.data.GetHashedKey()

	if key == nKey {
		if n.data.Expired() {
			return nil, errors.ErrKeyExpired
		} else {
			return n.data.Clone(), nil
		}
	} else if key < nKey {
		return getNode(n.left, k)
	} else {
		return getNode(n.right, k)
	}
}

// inOrderTraversal traverses the tree in-order (left, n, right).
func inOrderTraversal(n *node, r *[]*kv.KVPair) {
	if n != nil {
		inOrderTraversal(n.left, r)
		if !n.data.Expired() {
			*r = append(*r, n.data.Copy())
		}
		inOrderTraversal(n.right, r)
	}
}

func deleteNode(n *node, k []byte) *node {
	if n == nil {
		return nil
	}

	pKey := kv.HashKey(k)
	nKey := n.data.GetHashedKey()

	if pKey < nKey {
		n.left = deleteNode(n.left, k)
	} else if pKey > nKey {
		n.right = deleteNode(n.right, k)
	} else {
		// node to be deleted is found
		if n.left == nil {
			return n.right
		} else if n.right == nil {
			return n.left
		}
		// node has two children
		minRight := findMin(n.right)
		n.data = minRight.data.Copy()
		k, _ := minRight.data.GetKey()
		n.right = deleteNode(n.right, k)
	}

	return n
}

func findMin(n *node) *node {
	current := n
	for current.left != nil {
		current = current.left
	}
	return current
}
