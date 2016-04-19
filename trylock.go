package trylock

import (
	"sync"
	"sync/atomic"
	"unsafe"
)

const mutexLocked = 1 << iota

// Mutex is simple sync.Mutex + ability to try to Lock.
type Mutex struct {
	*sync.Mutex
}

// New returns new Mutex with initialized underlying sync.Mutex.
func New() *Mutex {
	return &Mutex{
		Mutex: &sync.Mutex{},
	}
}

// TryLock tries to Lock Mutex. It returns true in case of success, false
// otherwise.
func (mu *Mutex) TryLock() bool {
	state := (*int32)(unsafe.Pointer(mu.Mutex))
	return atomic.CompareAndSwapInt32(state, 0, mutexLocked)
}
