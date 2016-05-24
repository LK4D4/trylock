package trylock

import (
	"reflect"
	"testing"
)

func TestMutexLayout(t *testing.T) {
	sf := reflect.TypeOf((*Mutex)(nil)).Elem().FieldByIndex([]int{0, 0})
	if sf.Name != "state" {
		t.Fatal("sync.Mutex first field should have name state")
	}
	if sf.Offset != uintptr(0) {
		t.Fatal("sync.Mutex state field should have zero offset")
	}
	if sf.Type != reflect.TypeOf(int32(1)) {
		t.Fatal("sync.Mutex state field type should be int32")
	}
}

func TestTryLock(t *testing.T) {
	var mu Mutex
	if !mu.TryLock() {
		t.Fatal("mutex must be unlocked")
	}
	if mu.TryLock() {
		t.Fatal("mutex must be locked")
	}

	mu.Unlock()
	if !mu.TryLock() {
		t.Fatal("mutex must be unlocked")
	}
	if mu.TryLock() {
		t.Fatal("mutex must be locked")
	}

	mu.Unlock()
	mu.Lock()
	if mu.TryLock() {
		t.Fatal("mutex must be locked")
	}
	if mu.TryLock() {
		t.Fatal("mutex must be locked")
	}
	mu.Unlock()
}

func TestTryLockPointer(t *testing.T) {
	mu := &Mutex{}
	if !mu.TryLock() {
		t.Fatal("mutex must be unlocked")
	}
	if mu.TryLock() {
		t.Fatal("mutex must be locked")
	}

	mu.Unlock()
	if !mu.TryLock() {
		t.Fatal("mutex must be unlocked")
	}
	if mu.TryLock() {
		t.Fatal("mutex must be locked")
	}

	mu.Unlock()
	mu.Lock()
	if mu.TryLock() {
		t.Fatal("mutex must be locked")
	}
	if mu.TryLock() {
		t.Fatal("mutex must be locked")
	}
	mu.Unlock()
}

func TestRace(t *testing.T) {
	var mu Mutex
	var x int
	for i := 0; i < 1024; i++ {
		if i%2 == 0 {
			go func() {
				if mu.TryLock() {
					x++
					mu.Unlock()
				}
			}()
			continue
		}
		go func() {
			mu.Lock()
			x++
			mu.Unlock()
		}()
	}
}
