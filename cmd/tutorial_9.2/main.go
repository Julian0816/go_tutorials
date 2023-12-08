package main

import (
    "fmt"
    "time"
)

func main() {
    // The conveyor belt is set up to carry 5 toys.
    var conveyor = make(chan int, 5)

    // Your friend starts the conveyor belt to receive the toys.
    go startConveyorBelt(conveyor)

    // You wait and collect each toy that comes out of the conveyor belt.
    for toy := range conveyor {
        fmt.Println(toy) // Show the toy you've received.
        time.Sleep(time.Second * 1) // Take a 1-second break before getting the next toy.
    }
}

// This function puts toys on the conveyor belt.
func startConveyorBelt(conveyor chan int) {
    // After putting all the toys on the belt, you'll turn it off.
    defer close(conveyor)

    // Here are your toys, numbered from 0 to 4.
    for i := 0; i < 5; i++ {
        conveyor <- i // Place a toy on the conveyor belt.
    }

    // Announce that you're done putting toys on the belt.
    fmt.Println("Exiting process")
}
