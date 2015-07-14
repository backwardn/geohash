package geohash

import "testing"

// TestCase objects are generated from independent code to verify we get the
// same results. See testcases_test.go.
type TestCase struct {
	hashInt  uint64
	hash     string
	lat, lng float64
}

// Test we get the same string geohashes.
func TestEncode(t *testing.T) {
	for _, c := range testcases {
		hash := Encode(c.lat, c.lng)
		if c.hash != hash {
			t.Errorf("incorrect encode string result for (%v,%v): %s != %s",
				c.lat, c.lng, c.hash, hash)
		}
	}
}

// Test we get the same integer geohashes.
func TestEncodeInt(t *testing.T) {
	for _, c := range testcases {
		hashInt := EncodeInt(c.lat, c.lng)
		if c.hashInt != hashInt {
			t.Errorf("incorrect encode integer result for (%v,%v): %016x != %016x xor %016x",
				c.lat, c.lng, c.hashInt, hashInt, c.hashInt^hashInt)
		}
	}
}
