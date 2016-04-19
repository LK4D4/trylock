package trylock_test

import (
	"fmt"

	"github.com/LK4D4/trylock"
)

func Example() {
	mu := trylock.New()
	fmt.Println(mu.TryLock())
	fmt.Println(mu.TryLock())
	mu.Unlock()
	fmt.Println(mu.TryLock())
	// Output:
	// true
	// false
	// true
}
