package main

import (
	"fmt"
	"kvstore-db/kvstore"
)

func main() {
	store := kvstore.NewKVStore()

	store.Put([]byte("mario"), []byte("alfredo"))
	store.Put([]byte("nercia"), []byte("chale"))
	store.Put([]byte("elves"), []byte("junior"))
	store.Put([]byte("maria"), []byte("moiane"))

	value, err := store.Get([]byte("mario"))
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("Value found: ", string(value))

	value, err = store.Get([]byte("elves"))
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("Value found: ", string(value))

	value, err = store.Get([]byte("ana"))
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("Value found: ", string(value))
}
