package kvstore

import (
	"crypto/sha256"
	"errors"
	"fmt"
	"sync"
	"time"
)

type valueType struct {
	data       []byte
	expires    bool
	expiration time.Timer
}

type KVStore struct {
	mu   sync.RWMutex
	data map[string]valueType
}

func NewKVStore() *KVStore {
	return &KVStore{
		mu:   sync.RWMutex{},
		data: make(map[string]valueType),
	}
}

func (kv *KVStore) Put(key, value []byte, options ...any) {
	kv.mu.Lock()
	defer kv.mu.Unlock()

	item := valueType{data: value}
	kv.data[getHashedKey(key)] = item
}

func (kv *KVStore) Get(key []byte) ([]byte, error) {
	kv.mu.RLock()
	defer kv.mu.RUnlock()

	if value, exists := kv.data[getHashedKey(key)]; exists {
		return value.data, nil
	}

	return []byte{}, errors.New(`{"error": "key not found"}`)
}

func (kv *KVStore) Delete(key []byte) {
	kv.mu.Lock()
	defer kv.mu.Unlock()
	delete(kv.data, getHashedKey(key))
}

func (kv *KVStore) Exists(key []byte) bool {
	kv.mu.RLock()
	defer kv.mu.RUnlock()
	_, exists := kv.data[getHashedKey(key)]
	return exists
}

func getHashedKey(key []byte) string {
	hash := sha256.Sum256(key)
	hashStr := fmt.Sprintf("%x", hash)
	return hashStr
}
