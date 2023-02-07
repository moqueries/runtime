package moq_test

import (
	"testing"

	"moqueries.org/runtime/moq"
)

func TestFnString(t *testing.T) {
	// ASSEMBLE
	fn := func(string) error { return nil }

	// ACT
	out := moq.FnString(fn)

	// ASSERT
	expected := "moqueries.org/runtime/moq_test.TestFnString.func1"
	if out != expected {
		t.Errorf("got %s, want %s", out, expected)
	}
}
