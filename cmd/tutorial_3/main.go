package main

import (
	"errors"
	"fmt"
)

func main() {
	// The parameter is inforsed so I cannot pass an int as parameter
	var printValue string = "Hello, Julz!"
	printName(printValue)

	var numerator int = 11
	var denominator int = 2
	var result, remainder, err = intDivision(numerator, denominator)
	// if err != nil {
	// 	fmt.Println(err.Error())
	// } else if remainder == 0 {
	// 	fmt.Printf("The result of the integer division is %v", result)
	// } else {
	// fmt.Printf("The result of the integer division is %v with remainder %v", result, remainder)
	// }

	// Do the same as above but using a switch statement
	switch {
		case err != nil:
			fmt.Println(err.Error())
			//break <==> We dont need to write the breaks as they are implied
		case remainder == 0:
			fmt.Printf("The result of the integer division is %v", result)
		default:
			fmt.Printf("The result of the integer division is %v with remainder %v and ", result, remainder)
	}

	// This is a conditional switch statement with the remainder being the condition
	switch remainder {
		case 0:
			fmt.Printf("The division was exact")
		case 1,2:
			fmt.Printf("The division was close")
		default:
			fmt.Printf(" The division was not close")
	}
}

func printName(printValue string) {
	fmt.Println(printValue)
}

// Here you need to let go know what type of value you are returning in this case and int
//  func intDivision(numerator int, denominator int) int{
// 	var result int = numerator/denominator
// 	return result // 5
//  }

// building on the function above, this is how to return multiple values
func intDivision(numerator int, denominator int) (int, int, error) {
	var err error
	if denominator == 0 {
		err = errors.New("cannot divide by zero")
		return 0, 0, err
	}
	var result int = numerator / denominator
	var remainder int = numerator % denominator
	return result, remainder, err
}
