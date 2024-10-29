package kvstore

import (
	"crypto/sha256"
	"fmt"
	"sync"
	"time"
)

type valueType struct {
	data       []byte
	expires    bool
	expiration time.Time
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

func (kv *KVStore) Put(key, value []byte, options ...any) error {
	_key := getHashedKey(key)
	_item := valueType{data: value}

	// If options are passed, options[0] represents the time to live
	// of the key-value pair in the database.
	var duration time.Duration
	if len(options) > 0 {
		if du, ok := options[0].(time.Duration); ok {
			duration = du
		}
	}

	if duration > 0 {
		_item.expires = true
		// TODO: verify the logic of the duration of a key-value pair.
		// Does it compare to the time in which the user set?
		_item.expiration = time.Now().Add(duration)

		// TODO: make sure the following goroutine uses a fair interval (now is using every 5min)
		// to cleanup key-value pairs.
		go func() {
			ticker := time.NewTicker(time.Minute * 5)
			for {
				if _, ok := <-ticker.C; ok {
					kv.mu.Lock()
					if value, exists := kv.data[_key]; exists {
						if keyExpired(value.expiration) {
							delete(kv.data, _key)
							kv.mu.Unlock()
							return
						}
					}
					kv.mu.Unlock()
				}
			}
		}()
	}

	kv.mu.Lock()
	defer kv.mu.Unlock()

	// The following block make sure to update a non-expired key-value pair.
	if _, exists := kv.data[_key]; exists {
		if keyExpired(_item.expiration) {
			delete(kv.data, _key)
			return fmt.Errorf(`{"error": "'%s' key expired"}`, string(key))
		}
	}

	kv.data[_key] = _item
	return nil
}

func (kv *KVStore) Get(key []byte) ([]byte, error) {
	_key := getHashedKey(key)

	kv.mu.Lock()
	defer kv.mu.Unlock()

	if value, exists := kv.data[_key]; exists {
		if keyExpired(value.expiration) {
			delete(kv.data, _key)
			return []byte{}, fmt.Errorf(`{"error": "'%s' key expired"}`, string(key))
		}
		return value.data, nil
	}

	return []byte{}, fmt.Errorf(`{"error": "'%s' key not found"}`, string(key))
}

func (kv *KVStore) Delete(key []byte) {
	kv.mu.Lock()
	defer kv.mu.Unlock()
	delete(kv.data, getHashedKey(key))
}

func (kv *KVStore) Exists(key []byte) bool {
	_key := getHashedKey(key)

	kv.mu.Lock()
	defer kv.mu.Unlock()

	if value, exists := kv.data[_key]; exists {
		if keyExpired(value.expiration) {
			delete(kv.data, _key)
			return false
		}
		return true
	}

	return false
}

func (kv *KVStore) keyExists(key []byte) bool {
	_key := getHashedKey(key)

	if value, exists := kv.data[_key]; exists {
		if keyExpired(value.expiration) {
			delete(kv.data, _key)
			return false
		}
		return true
	}
	return false
}

func getHashedKey(key []byte) string {
	hash := sha256.Sum256(key)
	hashStr := fmt.Sprintf("%x", hash)
	return hashStr
}

func keyExpired(expiration time.Time) bool {
	return expiration.Compare(time.Now()) >= 0
}
