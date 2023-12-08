package main

import "fmt"

// Imagine you have a set of toy blocks, and you want to pass them through a slide (the 
// channel 'c') to your friend one by one. Each block has a number on it.


func main() {
    // Here's the slide you'll send the blocks down.
    var slide = make(chan int)

    // Tell your friend to start collecting the blocks from the end of the slide.
    go sendBlocksDownSlide(slide)

    // Now, wait and see each block as it comes down the slide.
    for block := range slide {
        fmt.Println(block)
    }
}

// You're ready to send blocks down the slide to your friend.
func sendBlocksDownSlide(slide chan int) {
    // Once you've sent all the blocks, you'll close the slide so your friend knows there are no more blocks coming.
    defer close(slide)

    // These are your blocks, numbered 0 to 4.
    for i := 0; i < 5; i++ {
        // Here goes a block down the slide!
        slide <- i
    }
}

// In this story:

// - The slide is the channel you're using to send the blocks (integers).
// - The sendBlocksDownSlide function is like you sending the blocks down the slide.
// - The defer close(slide) line is like you telling your friend that no more blocks will come after you send the last one.
// - The for block := range slide loop in main is like your friend watching the blocks come down the slide and counting them.
// - The slide <- i is you actually sending each block down the slide to your friend.
// - Just like how your friend knows to expect five blocks because you told them, the main function in the code expects to receive five integers over the channel. 
// Once all the integers are received and there are no more to come (close(slide)), the loop stops and the program ends.