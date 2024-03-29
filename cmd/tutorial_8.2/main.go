package main

import (
    "fmt"
    "time"
    "sync"
)

var wg = sync.WaitGroup{}
var dbData = []string{"id1", "id2", "id3", "id4", "id5"}

func main() {
    t0 := time.Now()
    for i := 0; i < 1000; i++ {
        wg.Add(1)
        go dbCall(i)
    }
    wg.Wait()
    fmt.Printf("\nTotal execution time: %v", time.Since(t0))
}

func dbCall(i int) {
    // Simulate DB call delay
    var delay float32 = 2000
    time.Sleep(time.Duration(delay) * time.Millisecond)
    wg.Done()
}
