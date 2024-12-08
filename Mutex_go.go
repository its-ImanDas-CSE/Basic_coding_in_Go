// A mutex is a lock used to protect shared data from being accessed simultaneously by multiple goroutines.
//It ensures that only one goroutine can access the data at a time, which prevents errors in a concurrent program.

// In Go, the line mu.Lock() is used to lock a mutex, and mu.Unlock() is used to unlock it. The main purpose of the mutex is to prevent concurrent access to shared data (in this case, counter) by multiple goroutines, ensuring that only one goroutine can modify counter at a time.
// The defer keyword in Go schedules a function (in this case, mu.Unlock()) to be executed after the surrounding function finishes, no matter what happens in the function. This is useful because it ensures that the mutex will be unlocked when the function is done, even if there’s an error or an early return in the function.

package main

import (
	"fmt"
	"sync"
)

// Shared counter variable
var counter int // global variable

// Mutex to control access to the counter
var mu sync.Mutex

// Function to safely increase the counter
func increment() {
	mu.Lock()         // Lock the mutex
	defer mu.Unlock() // Ensure the mutex is unlocked after the function finishes, This ensures that no goroutine holds the lock for too long, which could cause other goroutines to wait forever (a situation called "deadlock").

	counter++ // Safe increment
}

func main() {
	var wg sync.WaitGroup // We use a sync.WaitGroup to make sure the main program waits for all the goroutines to finish their work before printing the final value of the counter.

	// This for loop, Start 5 goroutines to increase the counter
	for i := 0; i < 5; i++ {
		wg.Add(1)   // wg is a sync.WaitGroup. .Add(1) tell the WaitGroup, "Hey, I just started one more task. Keep track of it.
		go func() { //This starts a goroutine. The go keyword tells Go to run the function concurrently.
			defer wg.Done() // wg.Done() tells the WaitGroup that this goroutine has finished its work. By using defer, we ensure that wg.Done() is called automatically when the goroutine exits
			increment()     // Increase the counter
		}()
	}
	/*1) A WaitGroup is used to wait for a collection of goroutines (concurrent tasks) to finish.
	2) You can think of it like a counter that keeps track of how many tasks are still running.
	3) The wg.Add(1) line is used to increase the counter of the WaitGroup by 1.
	4) It tells the WaitGroup that we are starting one more task.
	5) Since we are running multiple goroutines concurrently,
	6) we need to let the WaitGroup know how many goroutines we are expecting to finish before we move on.
	7) Each time we start a goroutine, we use wg.Add(1) to tell the WaitGroup, "Hey, I just started one more task. Keep track of it.
	8) " Later, when the task (goroutine) finishes, we use wg.Done() to reduce the counter.  */

	// Wait for all goroutines to finish
	wg.Wait()

	// Print the final value of counter
	fmt.Println("Final counter value:", counter)
}
