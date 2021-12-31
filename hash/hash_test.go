package hash_test

import (
	"testing"

	"github.com/myshkin5/moqueries/hash"
)

func TestNil(t *testing.T) {
	// ASSEMBLE

	// ACT
	h := hash.DeepHash(nil)

	// ASSERT
	// The value doesn't really matter. Just make sure it doesn't panic.
	if h != 0 {
		t.Errorf("wanted 0, got %d", h)
	}
}