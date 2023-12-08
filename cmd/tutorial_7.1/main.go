package main

import "fmt"



// 1) Creating an Array: In the 'main' function, you create an array called 'thing1' with 
// five elements of type 'float64'.

// 2) Printing Memory Address: You then print the memory address of 'thing1' using the 
// '&' operator, which gives you the address of the first element of the array.

// 3) Calling square Function: You pass a pointer to the 'thing1' array to the 'square'
// function. In Go, arrays are values, so if you pass an array to a function, it gets a copy. 
// But here, you are passing a pointer to the array, so 'square' can modify the original 
// array.

// 4) Modifying Array in square Function: Inside the 'square' function, 'thing2' is a 
// pointer to an array of five 'float64s'. It prints the memory address of 'thing2'
// which should be the same as 'thing1' because they point to the same data. The 
// function iterates over 'thing2', squaring each number in place.

// 5) Returning Modified Array: Finally, 'square' returns the dereferenced pointer to the
// modified array (now containing the squares of the original numbers).

// 6) Printing Results: Back in 'main', you print the 'result', which shows the squared 
// numbers. You also print the value of 'thing1' to demonstrate that 'thing1' was 
// modified by the 'square' function.

func main() {
    // Initialize an array of float64 with 5 elements.
    var thing1 = [5]float64{1, 2, 3, 4, 5}

    // Print the memory address of the thing1 array.
    fmt.Printf("\nThe memory location of the thing1 array is: %p", &thing1)

    // Pass a pointer to the thing1 array to the square function and store the result.
    var result [5]float64 = square(&thing1)

    // Print the resulting array after squaring its elements.
    fmt.Printf("\nThe result is: %v", result)

    // Print the updated thing1 array to show it has been modified.
    fmt.Printf("\nThe value of thing1 is: %v", thing1)
}

// Define the square function that takes a pointer to an array of float64s and returns an array of float64s.
func square(thing2 *[5]float64) [5]float64 {
    // Print the memory location of the thing2 array, which should be same as thing1.
    fmt.Printf("\nThe memory location of the thing2 array is: %p", thing2)

    // Iterate over the array and square each element.
    for i := range thing2 {
        thing2[i] = thing2[i] * thing2[i]
    }

    // Return the dereferenced and modified array.
    return *thing2
}
