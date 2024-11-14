package main

import (
	"fmt"
	"time"

	"github.com/imariom/nexusdb/pkg/bst"
	kv "github.com/imariom/nexusdb/pkg/kvpair"
)

func main() {
	bst := &bst.BST{}

	values := []kv.KVPair{
		kv.NewKVPair([]byte("userID123"), []byte("John Doe"), time.Second),
		kv.NewKVPair([]byte("sessionToken"), []byte("abc123xyz"), time.Minute*5),
		kv.NewKVPair([]byte("email"), []byte("jane.doe@example.com"), time.Hour*1),
		kv.NewKVPair([]byte("orderID456"), []byte("Order#789456"), time.Minute*30),
		kv.NewKVPair([]byte("productID"), []byte("Widget-X100"), time.Hour*2),
		kv.NewKVPair([]byte("imageData"), []byte{0xFF, 0xD8, 0xFF, 0xE0, 0x00, 0x10, 0x4A, 0x46, 0x49, 0x46}, time.Minute*20), // JPEG header byte,
		kv.NewKVPair([]byte("binaryData"), []byte{0x0A, 0x1B, 0x2C, 0x3D, 0x4E, 0x5F, 0x6A, 0x7B}, time.Second*45),            // Arbitrary binary sequence
	}

	for _, v := range values {
		bst.Insert(v)
	}

	result := bst.InOrder()
	for _, r := range result {
		key, err := r.GetKey()
		if err != nil {
			fmt.Println("Pair expired")
		}
		value, _ := r.GetValue()
		fmt.Println(string(key), string(value))
	}
}
