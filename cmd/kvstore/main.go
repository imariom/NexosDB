package main

import (
	"fmt"

	"github.com/imariom/NexusDB/kvstore"
)

func main() {
	store := kvstore.NewKVStore()

	store.Put([]byte("microsoft"), []byte("satya"))
	store.Put([]byte("apple"), []byte("cook"))
	store.Put([]byte("google"), []byte("pichai"))

	value, err := store.Get([]byte("microsoft"))
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("Value found: ", string(value))

	value, err = store.Get([]byte("apple"))
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("Value found: ", string(value))

	value, err = store.Get([]byte("amazon"))
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("Value found: ", string(value))
}
