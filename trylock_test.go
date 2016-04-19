package trylock

import "testing"

func TestTryLock(t *testing.T) {
	mu := New()
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
	mu := New()
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
