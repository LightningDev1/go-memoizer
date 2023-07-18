package memoizer

import (
	"testing"
	"time"
)

func TestMemoizer(t *testing.T) {
	t.Log("Testing memoizer...")

	functionCalls := 0

	function := func() (*int, error) {
		functionCalls++

		return &functionCalls, nil
	}

	memoizer := New(function, time.Second*2)

	first, _ := memoizer.Get()

	// The first call should be 1.
	if *first != 1 {
		t.Errorf("First call was not 1")
	}

	second, _ := memoizer.Get()

	// The second call should be cached.
	if *second != 1 {
		t.Errorf("Second call was not 1")
	}

	// Sleep for 3 seconds to invalidate the value.
	time.Sleep(time.Second * 3)

	third, _ := memoizer.Get()

	// The third call should be 2, since the value was invalidated.
	if *third != 2 {
		t.Errorf("Third call was not 2")
	}

	t.Log("Done testing memoizer.")
}
