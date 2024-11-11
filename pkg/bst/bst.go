package bst

import (
	"sync"

	kv "github.com/imariom/nexusdb/pkg/kvpair"
)

// Node[T] represents a node in a generic BST.
type node struct {
	data  kv.KVPair
	left  *node
	right *node
}

// BST represents the binary search tree with generic type T.
type BST struct {
	root *node
	mu   sync.RWMutex
}

// Insertion
func (bst *BST) Insert(data kv.KVPair) {
	bst.mu.Lock()
	defer bst.mu.Unlock()

	if bst.root == nil {
		bst.root = &node{
			data: kv.KVPair{Key: data.Key, Value: data.Value},
		}
	} else {
		inserNode(bst.root, data)
	}
}

func inserNode(node_ *node, data kv.KVPair) {
	if data.Key < node_.data.Key {
		node_.left = &node{data: data}
	}
}

// Search
// In-Order Traversal
// Deletion
