package memoizer

import (
	"sync"
	"time"
)

// Memoizer can cache a value for a certain amount of time.
// Used for caching a function return value.
type Memoizer[T any] struct {
	// The function to call to get the value.
	Function func() (*T, error)

	// The time after which the value should be invalidated.
	InvalidateAfter time.Duration

	// The time at which the value was last loaded.
	LastLoaded time.Time

	// The mutex to use for locking.
	Mutex sync.RWMutex

	// The value of the memoizer.
	Value *T
}

// New creates a new memoizer.
func New[T any](function func() (*T, error), invalidateAfter time.Duration) *Memoizer[T] {
	return &Memoizer[T]{
		Function:        function,
		InvalidateAfter: invalidateAfter,
	}
}

// InvalidateTime returns the time after which the value should be invalidated.
func (m *Memoizer[T]) InvalidateTime() time.Time {
	return m.LastLoaded.Add(m.InvalidateAfter)
}

// LoadValue loads the value from the function.
func (m *Memoizer[T]) LoadValue() (*T, error) {
	m.Mutex.Lock()
	defer m.Mutex.Unlock()

	value, err := m.Function()
	if err != nil {
		return nil, err
	}

	m.Value = value
	m.LastLoaded = time.Now()

	return m.Value, nil
}

// GetCacheValue gets the current cached value.
func (m *Memoizer[T]) GetCacheValue() *T {
	m.Mutex.RLock()
	defer m.Mutex.RUnlock()

	return m.Value
}

// Get gets the value of the memoizer.
// If the value is not cached, it will call the function to get the value.
func (m *Memoizer[T]) Get() (*T, error) {
	value := m.GetCacheValue()

	if value == nil || time.Now().After(m.InvalidateTime()) {
		return m.LoadValue()
	}

	return value, nil
}
