package kvpair

import (
	"testing"
	"time"

	"github.com/imariom/nexusdb/pkg/errors"
)

// Test case for the NewKVPair function
func TestNewKVPair(t *testing.T) {
	key := []byte("userID123")
	value := []byte("John Doe")
	ttl := time.Minute * 5

	kv := NewKVPair(key, value, ttl)
	if kv == nil {
		t.Fatalf("Expected a valid KVPair, got nil")
	}

	// Test key, value, and expiration time
	if string(kv.key) != "userID123" {
		t.Errorf("Expected key 'userID123', got %s", kv.key)
	}

	if string(kv.value) != "John Doe" {
		t.Errorf("Expected value 'John Doe', got %s", kv.value)
	}

	if kv.expiration.Before(time.Now()) {
		t.Errorf("Expected expiration time to be in the future, got %s", kv.expiration)
	}
}

// Test case for the Validate method
func TestValidate(t *testing.T) {
	key := []byte("userID123")
	value := []byte("John Doe")
	ttl := time.Minute * 5

	kv := NewKVPair(key, value, ttl)

	// Valid KVPair
	if err := kv.Validate(); err != nil {
		t.Errorf("Expected valid KVPair, got error: %v", err)
	}

	// Invalid KVPair (empty value)
	kv.value = nil
	if err := kv.Validate(); err != errors.ErrKeyNotValid {
		t.Errorf("Expected 'ErrKeyNotValid' error, got: %v", err)
	}

	// Expired KVPair
	kv = NewKVPair(key, value, ttl)
	kv.expiration = time.Now().Add(-time.Minute * 10) // Set expiration in the past
	if err := kv.Validate(); err != errors.ErrKeyExpired {
		t.Errorf("Expected 'ErrKeyExpired' error, got: %v", err)
	}
}

// Test case for the UpdateValue method
func TestUpdateValue(t *testing.T) {
	key := []byte("userID123")
	value := []byte("John Doe")
	ttl := time.Minute * 5

	kv := NewKVPair(key, value, ttl)

	// Update value successfully
	if err := kv.UpdateValue([]byte("Jane Doe")); err != nil {
		t.Errorf("Expected to update value successfully, got error: %v", err)
	}

	// Validate value update
	if string(kv.value) != "Jane Doe" {
		t.Errorf("Expected value 'Jane Doe', got %s", kv.value)
	}
}

// Test case for the UpdateTTL method
func TestUpdateTTL(t *testing.T) {
	key := []byte("userID123")
	value := []byte("John Doe")
	ttl := time.Minute * 5

	kv := NewKVPair(key, value, ttl)

	// Update TTL successfully
	newTTL := time.Minute * 10
	if err := kv.UpdateTTL(newTTL); err != nil {
		t.Errorf("Expected to update TTL successfully, got error: %v", err)
	}

	// Validate expiration update
	if kv.expiration.Before(time.Now()) {
		t.Errorf("Expected expiration time to be in the future, got %s", kv.expiration)
	}

	// Clear expiration
	if err := kv.UpdateTTL(0); err != nil {
		t.Errorf("Expected to clear TTL successfully, got error: %v", err)
	}
	if !kv.expiration.IsZero() {
		t.Errorf("Expected expiration to be zero, got %s", kv.expiration)
	}
}

// Test case for the Clone method
func TestClone(t *testing.T) {
	key := []byte("userID123")
	value := []byte("John Doe")
	ttl := time.Minute * 5

	kv := NewKVPair(key, value, ttl)

	clone, err := kv.Clone()
	if err != nil {
		t.Errorf("Expected successful clone, got error: %v", err)
	}

	if !kv.Equal(clone) {
		t.Errorf("Expected clone to be equal to the original KVPair")
	}
}

// Test case for the Move method
func TestMove(t *testing.T) {
	key := []byte("userID123")
	value := []byte("John Doe")
	ttl := time.Minute * 5

	kv := NewKVPair(key, value, ttl)

	// Move data to a new KVPair
	newKV, err := kv.Move()
	if err != nil {
		t.Errorf("Expected successful move, got error: %v", err)
	}

	// Validate original KVPair is invalid
	if kv.IsValid() {
		t.Errorf("Expected original KVPair to be invalid after move")
	}

	// Validate new KVPair is valid
	if !newKV.IsValid() {
		t.Errorf("Expected new KVPair to be valid after move")
	}
}

// Test case for the IsExpired method
func TestIsExpired(t *testing.T) {
	key := []byte("userID123")
	value := []byte("John Doe")
	ttl := time.Minute * 5

	kv := NewKVPair(key, value, ttl)

	// Before expiration
	if kv.IsExpired() {
		t.Errorf("Expected KVPair to not be expired, but it is")
	}

	// After expiration
	kv.expiration = time.Now().Add(-time.Minute * 10) // Set expiration in the past
	if !kv.IsExpired() {
		t.Errorf("Expected KVPair to be expired, but it is not")
	}
}

// Test case for the Key method
func TestKey(t *testing.T) {
	key := []byte("userID123")
	value := []byte("John Doe")
	ttl := time.Minute * 5

	kv := NewKVPair(key, value, ttl)

	// Retrieve key
	retrievedKey, err := kv.Key()
	if err != nil {
		t.Errorf("Expected to retrieve key successfully, got error: %v", err)
	}

	if string(retrievedKey) != string(kv.key) {
		t.Errorf("Expected retrieved key to be '%s', got '%s'", kv.key, retrievedKey)
	}
}

// Test case for the HashedKey method
func TestHashedKey(t *testing.T) {
	key := []byte("userID123")
	value := []byte("John Doe")
	ttl := time.Minute * 5

	kv := NewKVPair(key, value, ttl)

	// Retrieve hashed key
	hashedKey, err := kv.HashedKey()
	if err != nil {
		t.Errorf("Expected to retrieve hashed key successfully, got error: %v", err)
	}

	// Validate hashed key format
	if len(hashedKey) != 64 {
		t.Errorf("Expected SHA-256 hash to be 64 characters long, got: %s", hashedKey)
	}
}
