package main
import (
	"fmt"
)

// Pointers in Go are like directions to a house. Imagine you have a toy in your house. 
// Instead of carrying the toy everywhere, you just tell people where your house is. That's 
// what pointers do; they tell where the data is stored, like giving directions to your toy's 
// location. 🏠 ➡️🧸

// Pointers are data types that store memory locations rather than values like integers
// or floats for example

// To create a pointer I use the * syntax

// func main(){
//     // Create a pointer p of type *int32 and initialize it with memory for an int32.
//     var p *int32 = new(int32)
    
//     // Create an int32 variable i with default value 0.
//     var i int32
    
//     // Print the value that p points to (should be 0 as it's uninitialized).
//     fmt.Printf("The value p points to is: %v", *p)
    
//     // Print the value of i (should be 0 as it's uninitialized).
//     fmt.Printf("\nThe value if i is: %v", i)
    
//     // Make the pointer p point to the address of i.
//     p = &i
    
//     // Change the value of i to 1 by dereferencing p and assigning a new value.
//     *p = 1
    
//     // Print the value that p points to (should be 1 after the assignment).
//     fmt.Printf("\nThe value p points to is: %v", *p)
    
//     // Print the current value of i (should be 1 as we modified it through p).
//     fmt.Printf("\nThe value if i is: %v", i)
    
//     // Create another int32 variable k and initialize it with the value 2.
//     var k int32 = 2
    
//     // Assign the value of k to i, now i also becomes 2.
//     i = k
// }


func main() {
    // Create a slice of int32 containing the elements 1, 2, and 3.
    var slice = []int32{1, 2, 3}
    
    // Create a copy of the slice, but not the actual integers it points to.
    // sliceCopy now refers to the same underlying array as slice.
    var sliceCopy = slice
    
    // Change the third element of the sliceCopy (and thus the third element of slice) to 4.
    sliceCopy[2] = 4
    
    // Print the original slice - it will display [1 2 4] because sliceCopy and slice
    // point to the same underlying array and we modified the third element via sliceCopy.
    fmt.Println(slice)
    
    // Print the sliceCopy - it will also display [1 2 4] showing the change we made.
    fmt.Println(sliceCopy)
}
