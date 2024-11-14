package bst

import (
	"testing"
	"time"

	kv "github.com/imariom/nexusdb/pkg/kvpair"
)

func TestBST_InsertUpdateAndSearch(t *testing.T) {
	bst := &BST{}

	// Insert nodes
	bst.Insert(kv.NewKVPair([]byte("userID123"), []byte("John Doe"), time.Second*1))
	bst.Insert(kv.NewKVPair([]byte("sessionToken"), []byte("abc123xyz"), time.Minute*5))
	bst.Insert(kv.NewKVPair([]byte("email"), []byte("jane.doe@example.com"), time.Hour*1))
	bst.Insert(kv.NewKVPair([]byte("orderID456"), []byte("Order#789456"), time.Minute*30))
	bst.Insert(kv.NewKVPair([]byte("productID"), []byte("Widget-X100"), time.Hour*2))
	bst.Insert(kv.NewKVPair([]byte("imageData"), []byte{0xFF, 0xD8, 0xFF, 0xE0, 0x00, 0x10, 0x4A, 0x46, 0x49, 0x46}, time.Minute*20)) // JPEG header bytes
	bst.Insert(kv.NewKVPair([]byte("binaryData"), []byte{0x0A, 0x1B, 0x2C, 0x3D, 0x4E, 0x5F, 0x6A, 0x7B}, time.Second*45))            // Arbitrary binary sequence

	// Test search for inserted nodes
	if !bst.Search([]byte("userID123")) {
		t.Errorf("Expected to find 'John Doe' in BST")
	}
	if !bst.Search([]byte("sessionToken")) {
		t.Errorf("Expected to find 'abc123xyz' in BST")
	}
	if !bst.Search([]byte("email")) {
		t.Errorf("Expected to find 'jane.doe@example.com' in BST")
	}
	if !bst.Search([]byte("orderID456")) {
		t.Errorf("Expected to find 'Order#789456' in BST")
	}
	if !bst.Search([]byte("productID")) {
		t.Errorf("Expected to find 'Widget-X100' in BST")
	}
	if !bst.Search([]byte("imageData")) {
		t.Errorf("Expected to find JPEG header bytes in BST")
	}
	if !bst.Search([]byte("binaryData")) {
		t.Errorf("Expected to find arbitrary binary sequence in BST")
	}

	// Test search for a non-existent value
	if bst.Search([]byte("videoData")) {
		t.Errorf("Expected not to find videoData in BST")
	}
	if bst.Search([]byte("binary")) {
		t.Errorf("Expected not to find binary in BST")
	}

	// Test Getting values
	kv, err := bst.Get([]byte("sessionToken"))
	if err != nil {
		t.Errorf("Expected to get 'sessionToken'")
	}

	// Test updating nodes
	// Test search for updated nodes
}

// func TestBST_InOrderTraversal(t *testing.T) {
// 	bst := &BST{}

// 	values := []kv.KVPair{
// 		kvpair.NewKVPair([]byte("sessionToken"), []byte("abc123xyz"), time.Minute*5),
// 		kvpair.NewKVPair([]byte("email"), []byte("jane.doe@example.com"), time.Hour*1),
// 		kvpair.NewKVPair([]byte("orderID456"), []byte("Order#789456"), time.Minute*30),
// 		kvpair.NewKVPair([]byte("productID"), []byte("Widget-X100"), time.Hour*2),
// 	}

// 	for _, v := range values {
// 		bst.Insert(v)
// 	}

// 	// Expected in-order traversal order
// 	// expected := []kv.KVPair{
// 	// 	kvpair.NewKVPair([]byte("email"), []byte("jane.doe@example.com"), time.Hour*1),
// 	// 	kvpair.NewKVPair([]byte("orderID456"), []byte("Order#789456"), time.Minute*30),
// 	// 	kvpair.NewKVPair([]byte("productID"), []byte("Widget-X100"), time.Hour*2),
// 	// 	kvpair.NewKVPair([]byte("sessionToken"), []byte("abc123xyz"), time.Minute*5),
// 	// }
// 	result := bst.InOrder()
// 	for _, v := range result {
// 		t.Log(v)
// 	}

// 	// if !reflect.DeepEqual(result, expected) {
// 	// 	t.Errorf("Expected in-order traversal to be %v, got %v", expected, result)
// 	// }
// }

// func TestBST_Delete(t *testing.T) {
// 	bst := &BST{}
// 	values := []int{10, 20, 5, 15, 25}
// 	for _, v := range values {
// 		bst.Insert(v)
// 	}

// 	// Delete leaf node
// 	bst.Delete(25)
// 	if bst.Search(25) {
// 		t.Errorf("Expected not to find 25 in BST after deletion")
// 	}

// 	// Delete node with one child
// 	bst.Delete(20)
// 	if bst.Search(20) {
// 		t.Errorf("Expected not to find 20 in BST after deletion")
// 	}

// 	// Delete node with two children
// 	bst.Delete(10)
// 	if bst.Search(10) {
// 		t.Errorf("Expected not to find 10 in BST after deletion")
// 	}

// 	// Final in-order traversal after deletions
// 	expected := []int{5, 15}
// 	result := bst.InOrder()
// 	if !reflect.DeepEqual(result, expected) {
// 		t.Errorf("Expected in-order traversal to be %v, got %v", expected, result)
// 	}
// }

// func TestBST_Height(t *testing.T) {
// 	bst := &BST{}
// 	values := []int{10, 5, 15, 3, 7, 12, 18}
// 	for _, v := range values {
// 		bst.Insert(v)
// 	}

// 	expectedHeight := 3
// 	if bst.Height() != expectedHeight {
// 		t.Errorf("Expected height to be %d, got %d", expectedHeight, bst.Height())
// 	}
// }

// func TestBST_Size(t *testing.T) {
// 	bst := &BST{}
// 	values := []int{10, 20, 5, 15}
// 	for _, v := range values {
// 		bst.Insert(v)
// 	}

// 	expectedSize := 4
// 	if bst.Size() != expectedSize {
// 		t.Errorf("Expected size to be %d, got %d", expectedSize, bst.Size())
// 	}

// 	// Delete a node and test size again
// 	bst.Delete(15)
// 	expectedSize = 3
// 	if bst.Size() != expectedSize {
// 		t.Errorf("Expected size to be %d after deletion, got %d", expectedSize, bst.Size())
// 	}
// }
