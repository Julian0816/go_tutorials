// Main package where the program execution starts.
package main

// Importing necessary packages: fmt for formatting and printing,
// time for time-related operations, and sync for synchronization primitives.
import (
    "fmt"
    "time"
    "sync"
)

// Global variables:
// m is a read-write mutex that allows safe concurrent read & write access to shared resources.
// wg is a wait group that waits for a collection of goroutines to finish executing.
// dbData simulates a list of data to be processed (like database IDs).
// results hold the processed results.
var m = sync.RWMutex{}
var wg = sync.WaitGroup{}
var dbData = []string{"id1", "id2", "id3", "id4", "id5"}
var results = []string{}

// The main function is the entry point of the Go program.
func main() {
    // Record the current time to calculate the total execution time later.
    t0 := time.Now()
    // Loop over dbData and start a goroutine for each item.
    for i := 0; i < len(dbData); i++ {
        wg.Add(1)  // Increment the wait group counter.
        go dbCall(i)  // Start a goroutine for the database call.
    }
    wg.Wait()  // Wait for all goroutines to finish.
    // Print the total time taken to execute all goroutines.
    fmt.Printf("\nTotal execution time: %v", time.Since(t0))
    // Print the results collected from all goroutines.
    fmt.Printf("\nThe results are %v", results)
}

// dbCall simulates a database call for an item.
func dbCall(i int) {
    // Delay to simulate time taken for a database call.
    var delay float32 = 2000
    // Sleep for the duration of delay to simulate a database operation.
    time.Sleep(time.Duration(delay) * time.Millisecond)
    // Save the result of the database call.
    save(dbData[i])
    // Log the current state of the results.
    log()
    // Indicate to the wait group that the goroutine has finished.
    wg.Done()
}

// save adds a result to the results slice.
func save(result string) {
    m.Lock()  // Lock the mutex to prevent other goroutines from writing to results.
    results = append(results, result)  // Append the result to the results slice.
    m.Unlock()  // Unlock the mutex to allow other goroutines to write to results.
}

// log prints the current state of the results.
func log() {
    m.RLock()  // Lock the mutex for reading to prevent other goroutines from writing to results.
    // Print the current results.
    fmt.Printf("\nThe current results are: %v", results)
    m.RUnlock()  // Unlock the mutex to allow other goroutines to read or write to results.
}
