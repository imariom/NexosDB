package bst

import (
	"sync"

	kv "github.com/imariom/nexusdb/pkg/kvpair"
)

// Node[T] represents a node in a generic BST.
type node struct {
	data  *kv.KVPair
	left  *node
	right *node
}

// BST represents the binary search tree with generic type T.
type BST struct {
	root *node
	mu   sync.RWMutex
}

// Insertion
func (bst *BST) Insert(pair kv.KVPair) {
	bst.mu.Lock()
	defer bst.mu.Unlock()

	if bst.root == nil {
		bst.root = &node{
			data: pair.Clone(),
		}
	} else {
		inserNode(bst.root, pair)
	}
}

func inserNode(node_ *node, data kv.KVPair) {
	// When Key alredy exists update it's value
	// When key/value pair expired removed the KVPair
	if data.GetHashedKey() < node_.data.GetHashedKey() {
		if !node_.data.Expired() {
			node_.left = &node{
				data: data.Clone(),
			}
		}
	}
}

// Search
// In-Order Traversal
// Deletion
