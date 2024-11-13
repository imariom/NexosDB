// Package errors defines the error variables that may be returned
// during nexus operations.
package errors

import "errors"

// These errors can be returned when opening or calling methods on a DB.
var (
	// ErrDatabaseNotOpen is returned when a DB instance is accessed before it
	// is opened or after it is closed.
	ErrDatabaseNotOpen = errors.New("database not open")

	// ErrTimeout is returned when a database cannot obtain an exclusive lock
	// on the data file after the timeout passed to Open().
	ErrTimeout = errors.New("timeout")
)

// These errors can occur when creating, putting or deleting a key/value pair somewhere.
var (
	// ErrKeyNotFound is returned when trying to access a key that has
	// not been created yet.
	ErrKeyNotFound = errors.New("key not found")

	// ErrKeyExpired is returned when trying to access a key that has
	// expired.
	ErrKeyExpired = errors.New("key expired")

	// ErrKeyRequired is returned when inserting a zero-length key.
	ErrKeyRequired = errors.New("key required")

	// ErrKeyTooLarge is returned when inserting a key that is larger than MaxKeySize.
	ErrKeyTooLarge = errors.New("key too large")

	// ErrValueTooLarge is returned when inserting a value that is larger than MaxValueSize.
	ErrValueTooLarge = errors.New("value too large")
)
