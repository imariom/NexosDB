// Package kvpair defines the KVPair type that will be used
// to store and retrieve data in nexus operations.
package kvpair

import (
	"crypto/sha256"
	"fmt"
	"time"
	"unsafe"

	"github.com/imariom/nexusdb/pkg/errors"
)

// KVPair represents a pair of key/value to be stored somewhere.
type KVPair struct {
	key        []byte
	value      []byte
	expiration time.Time
	updatedAt  time.Time
}

// NewKVPair constructs a KVPair object that represent a key/value pair somewhere.
func NewKVPair(key, value []byte, duration time.Duration) *KVPair {
	var expire time.Time
	if duration > 0 {
		expire = time.Now().Add(duration)
	}

	return &KVPair{
		key:        key,
		value:      value,
		expiration: expire,
		updatedAt:  time.Now(),
	}
}

// GetKey returns the original non-hashed key of the pair.
func (kv *KVPair) GetKey() ([]byte, error) {
	if keyExpired(kv.expiration) {
		return nil, errors.ErrKeyExpired
	}

	var key []byte
	key = append(key, kv.key...)

	return key, nil
}

// GetHashedKey transform keys into a fixed-length SHA-256 hash.
func (kv *KVPair) GetHashedKey() string {
	return getHashedKey(kv.key)
}

// GetValue get a copy of the value of the non-expired KVPair.
func (kv *KVPair) GetValue() ([]byte, error) {
	if keyExpired(kv.expiration) {
		return nil, errors.ErrKeyExpired
	}

	var value []byte
	value = append(value, kv.value...)

	return value, nil
}

// SetExpiration sets the expiration time for the key-value pair based on a duration.
// It also updates the updatedAt field to the current time.
//
// Parameters:
//   - duration (time.Duration): The duration from now after which the key-value pair expires.
//
// Example:
//
//	kv.SetExpiration(10 * time.Minute)  // Expires 10 minutes from now.
func (kv *KVPair) SetExpiration(duration time.Duration) error {
	if keyExpired(kv.expiration) {
		return errors.ErrKeyExpired
	}

	kv.expiration = time.Now().Add(duration)
	kv.updatedAt = time.Now()

	return nil
}

// SetValue sets the value of the KVPair and updates the updatedAt timestamp.
//
// Parameters:
//   - newValue ([]byte): The new value to set for the KVPair.
//
// Example:
//
//	kv.SetValue([]byte("newValue"))
func (kv *KVPair) SetValue(newValue []byte) error {
	if keyExpired(kv.expiration) {
		return errors.ErrKeyExpired
	}

	var v []byte
	v = append(v, newValue...)
	kv.value = v
	kv.updatedAt = time.Now()

	return nil
}

// Size returns the total size of the key, value, expiration and updatedAt fields in bytes.
//
// Returns:
//   - int: The combined length of key and value fields in bytes.
//
// Example:
//
//	size := kv.Size()
func (kv *KVPair) Size() int {
	var t time.Time
	s := unsafe.Sizeof(t)

	return len(kv.key) + len(kv.value) + int(s)*2
}

// Expired checks whether the current key/value pair has expired.
func (kv *KVPair) Expired() bool {
	return keyExpired(kv.expiration)
}

// getHashedKey generates a SHA-256 hash of a given key and returns it as a hexadecimal string.
//
// Parameters:
//   - key ([]byte): The input key to be hashed.
//
// Returns:
//   - string: The SHA-256 hash of the key in hexadecimal format.
//
// Example:
//
//	key := []byte("my_key")
//	hashedKey := getHashedKey(key)
//	fmt.Println(hashedKey)  // Outputs a hexadecimal string representing the hash of "my_key"
//
// Notes:
//
//	This function can be used to transform keys into a fixed-length hash, suitable for use
//	in a key-value storage engine where a consistent and unique key representation is needed.
func getHashedKey(key []byte) string {
	hash := sha256.Sum256(key)
	hashstr := fmt.Sprintf("%x", hash)
	return hashstr
}

// keyExpired checks if a given expiration time has passed.
//
// Parameters:
//   - expiration (time.Time): The time at which the key is set to expire.
//
// Returns:
//   - bool: Returns true if the current time is past the expiration time (key has expired),
//     and false otherwise.
//
// Example:
//
//	expTime := time.Now().Add(-time.Minute)  // Expired one minute ago
//	expired := keyExpired(expTime)
//	fmt.Println(expired)  // Outputs: true
func keyExpired(expiration time.Time) bool {
	return expiration.Before(time.Now())
}
