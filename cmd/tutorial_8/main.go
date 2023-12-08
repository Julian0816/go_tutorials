package main

import (
    "fmt"
    "time"
    "sync"
)

// Go Routines
// This is a way to launch multiple functions and have them execute concurrently

// IMPORTANT NOTE:
// Concurrency != Parallel execution
// Concurrency meand that I have multiple tasks in progress at the same time



var m sync.Mutex  // Create a mutex to prevent data races when writing to the results slice.
var wg sync.WaitGroup  // Create a wait group to manage synchronization between goroutines.
var dbData = []string{"id1", "id2", "id3", "id4", "id5"}  // Simulated database data.
var results = []string{}  // Slice to hold the results.

func main() {
    t0 := time.Now()  // Record the start time of the operation.

    for i := 0; i < len(dbData); i++ {
        wg.Add(1)  // Tell the wait group to expect another goroutine (another task).
        go dbCall(i)  // Start a new goroutine for the database call.
    }

    wg.Wait()  // Wait for all goroutines to finish their tasks.
    fmt.Printf("\nTotal execution time: %v", time.Since(t0))  // Print the total time taken to execute all goroutines.
    fmt.Printf("\nThe results are %v", results)  // Print the results collected from all goroutines.
}

// dbCall simulates a database call.
func dbCall(i int) {
    var delay float32 = 2000  // Simulate a delay, e.g., a database operation taking time.
    time.Sleep(time.Duration(delay) * time.Millisecond)  // Actually wait for the simulated delay.
    fmt.Println("The result from the database is:", dbData[i])  // Print a message with the result.

    m.Lock()  // Lock the mutex to ensure exclusive access to the results slice.
    results = append(results, dbData[i])  // Append the result to the slice.
    m.Unlock()  // Unlock the mutex so other goroutines can access results.

    wg.Done()  // Notify the wait group that this goroutine's task is complete.
}
