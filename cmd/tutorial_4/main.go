package main
import (
	"fmt"
)

func main(){

	// []Arrays
    // -------- 
    // Fixed length
    // Same Type
    // Indexable
    // Contiguous in memory

	// var intArr [3]int32 // Fixed lenght, same type, indexable [0,0,0]
	// intArr[1] = 123
	// fmt.Println(intArr[0])
	// fmt.Println(intArr[1:3])

	// fmt.Println(&intArr[0]) // Print out the memory location using &
	// fmt.Println(&intArr[1]) // Contiguous in memory
	// fmt.Println(&intArr[2])

	// intArr := [...]int32{1,2,3} // this [...] will imply the length of the array
	// fmt.Println(intArr)


	intArr := [3]int32{1,2,3} // You can initialize array when you define it
	fmt.Println(intArr)

	var intSlice []int32 = []int32{4,5,6} // A slice is like an array but you can append values to it
	fmt.Printf("The lenght is %v with capacity %v", len(intSlice), cap(intSlice))
	intSlice = append(intSlice, 7) // takes a slice as the first argument and what you want to append as the second
	fmt.Printf("\nThe lenght is %v with capacity %v\n", len(intSlice), cap(intSlice))
	// The lenght is 3 with capacity 3
	// The lenght is 4 with capacity 6

	// fmt.Println(intSlice[4]) // This is out of range
	// Ask Joe to explain this a bit better
	// [4 5 6]
	// [4 5 6 7]

	// You can also add multiple values to the slice wit the spread operator
	var intSlice2 []int32 = []int32{8,9}
	intSlice = append(intSlice, intSlice2...)
	fmt.Println(intSlice)

	// Specify the length of the slice as well as the capacity
	var intSlice3 []int32 = make([]int32, 3, 8)
	fmt.Println(intSlice3)

	// ----------------------------------------------------------------

	//Map
	var myMap map[string]uint8 = make(map[string]uint8)
	fmt.Println(myMap)

	var myMap2 = map[string]uint8{"Adam":23, "Sarah":45}
	fmt.Println(myMap2["Adam"]) // Get the age of Adam like this
	fmt.Println(myMap2["Jason"]) // if I try to get the value of a key that doesnt exist I get the default value of that type, in this case 0
	// Map will always return something, even if the key does not exists

	// Delete an element from the map
	// delete(myMap2, "Adam") // Takes the map name and the key (Comment this when studying)

	var age, ok = myMap2["Jason"]
	// Maps in go return a second optional value which is a boolean
	// this will return true if the value is in the map
	// and false if it is not
	if ok {
		fmt.Printf("The age is %v", age)
	} else {
		fmt.Println("Invalid Name")
	}

	// loops
	for name:= range myMap2{
		fmt.Printf("Name: %v\n", name ) // The order is not preserved
	}

	// if we want to also return the values we can do it like this
	for name2, age:= range myMap2{
		fmt.Printf("Name: %v, Age: %v \n", name2, age)
	}

	// Similarly we can iterate over arrays and slices like this
	for i, v := range intArr {
		fmt.Printf("Index: %v, Value: %v \n", i, v)
	}

	// This is a while loop in go
	// It will continue until i is greater or equal than 10
	var i int = 0
	for i<10{
		fmt.Println(i)
		i = i + 1
	}

	// I can omit the condition and instead use a break statement
	var j int = 0
	for {
		if j >= 10{
			break
		}
		fmt.Println(j)
		j = j + 1
	}

	// The same thing can be achieved with this syntax
	// initalization, condition, post
	for x:=0; x<10; x++ {
		fmt.Println(x)
	}

	// i++ Increments by 1
	// i += 10 Increments by 10
	// i *= 10 Multiply by 10

	// i-- Decrements by 1
	// i -= 10 Decrements by 10
	// i /= 10 Divides by 10

}