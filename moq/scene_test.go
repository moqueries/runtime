package moq_test

import (
	"testing"

	"moqueries.org/runtime/moq"
)

type mockMoq struct {
	resetCalled                 int
	assertExpectationsMetCalled int
}

func (m *mockMoq) Reset() {
	m.resetCalled++
}

func (m *mockMoq) AssertExpectationsMet() {
	m.assertExpectationsMetCalled++
}

func TestScene(t *testing.T) {
	var (
		moq1 *mockMoq
		moq2 *mockMoq
		moqT *mockT

		testScene *moq.Scene
	)

	beforeEach := func(t *testing.T) {
		t.Helper()

		moq1 = &mockMoq{}
		moq2 = &mockMoq{}
		moqT = &mockT{}

		testScene = moq.NewScene(moqT)
		testScene.AddMoq(moq1)
		testScene.AddMoq(moq2)
	}

	t.Run("resets all of its moqs", func(t *testing.T) {
		// ASSEMBLE
		beforeEach(t)

		// ACT
		testScene.Reset()

		// ASSERT
		if moq1.resetCalled != 1 {
			t.Errorf("got %d reset calls, want 1", moq1.resetCalled)
		}
		if moq2.resetCalled != 1 {
			t.Errorf("got %d reset calls, want 1", moq2.resetCalled)
		}
	})

	t.Run("asserts all of its moqs meet expectations", func(t *testing.T) {
		// ASSEMBLE
		beforeEach(t)

		// ACT
		testScene.AssertExpectationsMet()

		// ASSERT
		if moqT.helperCalled != 1 {
			t.Errorf("got %d helper calls, want 1", moqT.helperCalled)
		}
		if moq1.assertExpectationsMetCalled != 1 {
			t.Errorf("got %d assert expectations met calls, want 1",
				moq1.assertExpectationsMetCalled)
		}
		if moq2.assertExpectationsMetCalled != 1 {
			t.Errorf("got %d assert expectations met calls, want 1",
				moq2.assertExpectationsMetCalled)
		}
	})

	t.Run("returns the same MoqT it is given", func(t *testing.T) {
		// ASSEMBLE
		beforeEach(t)

		// ACT
		actualMoqT := testScene.T

		// ASSERT
		if actualMoqT != moqT {
			t.Errorf("got %#v, want %#v", actualMoqT, moqT)
		}
	})
}
