package main

import (
    "fmt"
    "math/rand"
    "time"
)

const MAX_CHICKEN_PRICE = 10
const MAX_TOFU_PRICE = 5

func main() {
    // Imagine these two channels are like phone lines to two different departments in a store.
    var chickenChannel = make(chan string)
    var tofuChannel = make(chan string)

    // Here's a list of three stores you're going to check.
    websites := []string{"walmart.com", "costco.com", "wholefoods.com"}

    // You're asking your store teams to check for chicken and tofu prices.
    // They'll use their phone lines (channels) to tell you if they find a good deal.
    for i := range websites {
        go checkChickenPrices(websites[i], chickenChannel)
        go checkTofuPrices(websites[i], tofuChannel)
    }

    // You're ready to listen to any good deals your teams find.
    sendMessage(chickenChannel, tofuChannel)
}

// This function listens for any good chicken or tofu deals and lets you know.
func sendMessage(chickenChannel, tofuChannel chan string) {
    // You're waiting to hear from either department.
    select {
    case website := <-chickenChannel: // Chicken department found a deal!
        fmt.Printf("\nText Sent: Found deal on chicken at %v.", website)
    case website := <-tofuChannel: // Tofu department found a deal!
        fmt.Printf("\nEmail Sent: Found deal on tofu at %v.", website)
    }
}

// Your chicken department is checking prices.
func checkChickenPrices(website string, chickenChannel chan string) {
    for {
        time.Sleep(time.Second * 1) // Taking a moment between checks.
        var chickenPrice = rand.Float32() * 20 // Checking the chicken price.
        if chickenPrice <= MAX_CHICKEN_PRICE { // If there's a good deal...
            chickenChannel <- website // Call you on the chicken phone line!
            break // Done checking.
        }
    }
}

// Your tofu department is doing the same thing.
func checkTofuPrices(website string, c chan string) {
    for {
        time.Sleep(time.Second * 1) // Taking a moment between checks.
        var tofu_price = rand.Float32() * 20 // Checking the tofu price.
        if tofu_price <= MAX_TOFU_PRICE { // Found a good deal!
            c <- website // Call you on the tofu phone line!
            break // Done checking.
        }
    }
}
