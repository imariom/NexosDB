package kvpair

import "testing"

func TestKVPairCreation(t *testing.T) {
	pair := NewKVPair([]byte("lorem"), []byte("The brown fox"), 0)

	// check key
	key, err := pair.GetKey()
	if err != nil {
		t.Error("Not able to GetKey:", err.Error())
	} else {
		if string(key) != "lorem" {
			t.Errorf("Expected key 'lorem', found '%s'\n", string(key))
		}
	}

	// check value
	val, err := pair.GetValue()
	if err != nil {
		t.Error("Not able to GetValue:", err.Error())
	} else {
		if string(val) != "The brown fox" {
			t.Errorf("Expected value 'The brown fox', found '%s'\n", string(val))
		}
	}
}
