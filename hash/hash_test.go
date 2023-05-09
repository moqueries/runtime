package hash_test

import (
	"testing"

	"moqueries.org/runtime/hash"
)

func TestNil(t *testing.T) {
	// ASSEMBLE

	// ACT
	//nolint:ifshort // short syntax here blurs ACT/ASSERT
	h := hash.DeepHash(nil)

	// ASSERT
	// The value doesn't really matter. Just make sure it doesn't panic.
	expect := hash.Hash(14695981039346656037)
	if h != expect {
		t.Errorf("got %d, wanted %d", h, expect)
	}
}

type t1 struct {
	a string
	b int
}

type t2 struct {
	a string
	b int
}

func TestTypes(t *testing.T) {
	// ASSEMBLE
	v1 := t1{
		a: "hi",
		b: 42,
	}
	v2 := t2{
		a: "hi",
		b: 42,
	}

	// ACT
	h1 := hash.DeepHash(v1)

	// ASSERT
	h2 := hash.DeepHash(v2)

	if h1 != h2 {
		t.Errorf("got different hashes (%d != %d), want same hashes", h1, h2)
	}

	set := make(map[hash.Hash]struct{})
	set[h1] = struct{}{}
	if _, ok := set[h2]; !ok {
		t.Errorf("got different set keys, want same keys")
	}
}

func TestAnonymousTypes(t *testing.T) {
	// ASSEMBLE
	v1 := t1{
		a: "hi",
		b: 42,
	}
	v2 := struct {
		x string
		y int
	}{
		x: "hi",
		y: 42,
	}

	// ACT
	h1 := hash.DeepHash(v1)

	// ASSERT
	h2 := hash.DeepHash(v2)

	if h1 != h2 {
		t.Errorf("got different hashes (%d != %d), want same hashes", h1, h2)
	}

	set := make(map[hash.Hash]struct{})
	set[h1] = struct{}{}
	if _, ok := set[h2]; !ok {
		t.Errorf("got different set keys, want same keys")
	}
}
