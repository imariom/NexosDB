// Package kvpair defines the KVPair type that will be used
// to store and retrieve key/value pair data in nexus operations.
package kvpair

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"time"

	"github.com/imariom/nexusdb/pkg/errors"
)

// KVPair represents a key-value pair with metadata, including expiration and update timestamps.
type KVPair struct {
	// key is the unique identifier for this key-value pair.
	// It is a byte slice to accommodate any kind of data, including strings, integers, serialized objects, or even binary data
	key []byte

	// value is the data associated with the key.
	// It is stored as a byte slice, allowing flexibility to store any kind of data, including text and binary formats.
	value []byte

	// expiration indicates the time at which this key-value pair is set to expire.
	// After this time, the pair should be considered invalid and may be subject to removal.
	expiration time.Time

	// updatedAt records the most recent time at which the key-value pair was modified.
	// This timestamp is useful for tracking changes and implementing caching or consistency mechanisms.
	updatedAt time.Time
}

// NewKVPair creates and returns a new KVPair with the provided key and value.
// If a ttl (Time-To-Live) duration is specified and greater than zero, expiration is set accordingly.
// Otherwise, expiration is set to a zero value, indicating no expiration.
func NewKVPair(key, value []byte, ttl time.Duration) *KVPair {
	var exp time.Time
	if ttl > 0 {
		exp = time.Now().Add(ttl)
	}

	return &KVPair{
		key:        key,
		value:      value,
		expiration: exp,
		updatedAt:  time.Now(),
	}
}

// UpdateValue updates the value of the KVPair and refreshes the updateAt timestamp.
func (kv *KVPair) UpdateValue(newValue []byte) error {
	if err := kv.Validate(); err != nil {
		return err
	}

	kv.value = newValue
	kv.updatedAt = time.Now()
	return nil
}

// UpdateTTL sets a new TTL (Time-To-Live) for the KVPair, updating the expiration time.
// If ttl is zero or less, the expiration is cleared, making the KVPair non-expiring.
func (kv *KVPair) UpdateTTL(ttl time.Duration) error {
	if err := kv.Validate(); err != nil {
		return err
	}

	if ttl > 0 {
		kv.expiration = time.Now().Add(ttl)
	} else {
		// Reset expiration to zero value, meaning no expiration.
		kv.expiration = time.Time{}
	}
	kv.updatedAt = time.Now()

	return nil
}

// Key returns a copy of the key to prevent external modification.
func (kv *KVPair) Key() ([]byte, error) {
	if err := kv.Validate(); err != nil {
		return nil, err
	}

	return append([]byte(nil), kv.key...), nil
}

// HashedKey transform keys into a fixed-length SHA-256 hash.
func (kv *KVPair) HashedKey() (string, error) {
	if err := kv.Validate(); err != nil {
		return "", err
	}

	return getHashedKey(kv.key), nil
}

// Value returns a copy of the value to prevent external modification.
func (kv *KVPair) Value() ([]byte, error) {
	if err := kv.Validate(); err != nil {
		return nil, err
	}

	return append([]byte(nil), kv.value...), nil
}

// UpdatedAt returns the last update time of the KVPair.
func (kv *KVPair) UpdatedAt() (time.Time, error) {
	if err := kv.Validate(); err != nil {
		return time.Time{}, err
	}
	return kv.updatedAt, nil
}

// Expiration returns the expiration time of the KVPair.
func (kv *KVPair) Expiration() (time.Time, error) {
	if err := kv.Validate(); err != nil {
		return time.Time{}, err
	}
	return kv.expiration, nil
}

// Clone creates a new KVPair with the same key, value, expiration, and updatedAt timestamps as the current instance.
// This ensures a deep copy of the key and value slices to prevent unintended mutations.
func (kv *KVPair) Clone() (*KVPair, error) {
	if err := kv.Validate(); err != nil {
		return nil, err
	}

	nk := append([]byte(nil), kv.key...)
	nv := append([]byte(nil), kv.value...)

	return &KVPair{
		key:        nk,
		value:      nv,
		expiration: kv.expiration,
		updatedAt:  kv.updatedAt,
	}, nil
}

// Move transfers all data from the current KVPair to the target KVPair,
// leaving the current KVPair in an invalid state.
func (kv *KVPair) Move() (*KVPair, error) {
	if err := kv.Validate(); err != nil {
		return nil, err
	}

	// Move all fields to the target
	tmp := &KVPair{
		key:        kv.key,
		value:      kv.value,
		expiration: kv.expiration,
		updatedAt:  kv.updatedAt,
	}

	// Invalidate the current KVPair by setting fields to zero values
	kv.key = nil
	kv.value = nil
	kv.expiration = time.Time{}
	kv.expiration = time.Time{}

	return tmp, nil
}

// IsExpired checks if the KVPair has expired. If the expiration time is zero, it is considered non-expiring.
func (kv *KVPair) IsExpired() bool {
	if kv.expiration.IsZero() {
		return false
	}
	return time.Now().After(kv.expiration)
}

// IsValid checks if the current KVPair is valid.
// A KVPair is considered valid if:
// - The key and value are non-nil and non-empty.
// - The expiration is either unset or set to a future time.
func (kv *KVPair) IsValid() bool {
	if len(kv.key) == 0 || len(kv.value) == 0 {
		return false
	}
	// if !kv.expiration.IsZero() && kv.expiration.Before(time.Now()) {
	// 	return false
	// }
	return true
}

// Validate checks if the KVPair is valid and not expired.
//
// This method performs two checks on the KVPair:
// 1. It verifies that the KVPair is valid by calling the IsValid method. If the key or value is empty, the pair is considered invalid, and an error is returned.
// 2. It checks if the KVPair has expired by calling the IsExpired method. If the expiration time is in the past, an error is returned.
//
// If both checks pass, the method returns nil, indicating that the KVPair is valid and not expired.
//
// Returns:
//   - nil: if the KVPair is valid and not expired.
//   - errors.ErrKeyNotValid: if the KVPair is invalid (i.e., key or value is empty).
//   - errors.ErrKeyExpired: if the KVPair has expired.
func (kv *KVPair) Validate() error {
	if !kv.IsValid() {
		return errors.ErrKeyNotValid
	}

	if kv.IsExpired() {
		return errors.ErrKeyExpired
	}

	return nil
}

// Equal checks if two KVPairs have the same key, value, expiration, and update times.
func (kv *KVPair) Equal(other *KVPair) bool {
	return bytes.Equal(kv.key, other.key) &&
		bytes.Equal(kv.value, other.value) &&
		kv.expiration.Equal(other.expiration) &&
		kv.updatedAt.Equal(other.updatedAt)
}

// HashedKey transform keys into a fixed-length SHA-256 hash.
func HashKey(key []byte) string {
	return getHashedKey(key)
}

// getHashedKey generates a SHA-256 hash of a given key and returns it as a hexadecimal string.
func getHashedKey(key []byte) string {
	hash := sha256.Sum256(key)
	hashstr := fmt.Sprintf("%x", hash)
	return hashstr
}
