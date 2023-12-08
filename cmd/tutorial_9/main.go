package main

import "fmt"

// Channels are a way to enable goroutines to pass around information
// Channels => Hold data
//          => Thread Safe
//          => Listen for data

// func main() {
// 	var c = make(chan int)
// 	c <- 1                          This is how not to use channels
// 	var i = <- c
// 	fmt.Println(i)
// }



// Define the main function which is the entry point of a Go program.
func main() {
    // Create a new channel 'c' that can transmit integers.
    var c = make(chan int)

    // Start a new goroutine that runs the 'process' function with the channel 'c' as its argument.
    go process(c)

    // Read an integer from channel 'c' and print it to the console.
    // This line will wait until an integer is sent on the channel.
    fmt.Println(<-c)
}

// Define a function named 'process' that accepts a channel 'c' for integers.
func process(c chan int) {
    // Send the integer 123 into the channel 'c'.
    // The main function's goroutine will receive this value.
    c <- 123
}



// Absolutely, let's make an analogy with the code you gave me:

// Imagine you have a basket of apples (the channel `c`) that you want to share with your friend (the `process` function). You can only give your friend one apple at a time, and your friend will only take one apple at a time from the basket.

// ```go
// func main() {
//     // You have a basket that can hold one apple at a time.
//     var basket = make(chan int)

//     // You ask your friend to get ready to take an apple from the basket.
//     go giveAppleToFriend(basket)

//     // You wait and watch your friend take the apple from the basket.
//     fmt.Println(<-basket)
// }

// // Your friend is ready to take an apple.
// func giveAppleToFriend(basket chan int) {
//     // Your friend takes an apple with the number 123 on it.
//     basket <- 123
// }
// ```

// In this story:
// - The `basket` (the channel `c`) is a way for you to send something (like an apple) to your friend.
// - The `giveAppleToFriend` function is like your friend taking an apple from the basket.
// - The `main` function is like you, waiting to see your friend take the apple.
// - The line `basket <- 123` is like putting an apple numbered 123 into the basket.
// - The line `<-basket` in `main` is you watching your friend take the apple out of the basket.

// Just like in the code, you have to wait for your friend to take the apple before you can continue playing. In Go, your program waits at `<-basket` until something is sent through the channel.