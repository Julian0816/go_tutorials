package main
import (
	"fmt"
	"strings"
)

func main(){
	var myString = "résumé" // keep an eye on the é character (this is an underlying array
	// of bites which represents the uft8 encoding of the full string which looks like this)

	// r [01110010]
	// é [11000011, 10101001]
	// s [01110011]
	// u [01110101]
	// m [01101101]
	// é [11000011, 10101001]


	var indexed = myString[0]
	fmt.Printf("%v, %T \n", indexed, indexed) // This returns this: 114, uint8 (Wired right?)
	for i, v := range myString { // is decoding é [11000011, 10101001] to 233 instead of 195 which is only the first part of the
		fmt.Println(i,v)
	}
	fmt.Printf("\nThe length of 'myString' is %v\n", len(myString))
	// Be careful with this because it returns the length of bites, not the characters

	// An easier way to deal with iterating and indexing strings is to cast them to an array of runes
	// rather than dealing with the underlying bite array of the string

	var myString2 = []rune("résumé")
	var indexed2 = myString2[1]
	fmt.Printf("%v, %T \n", indexed2, indexed2)
	for i, v := range myString2 {
		fmt.Println(i,v)
	}
	fmt.Printf("\nThe length of 'myString' is %v", len(myString2))

	// I can declare a rune type
	var myRune = 'a' 
	fmt.Printf("\nmyRune = %v", myRune) // This give me the unicode representation of 'a' which is 97

	// Concatenate strings
	var strSlice = []string{"j","u","l","i","a","n"}
	var catStr = ""
	for i := range strSlice{
		catStr += strSlice[i] // This is inefficient because I am creating a string every time, instead import "strings" package
	}
	fmt.Printf("\n%v", catStr)
	// String are unmutable in go, so once is created I cannot change it => catStr[0] = 'a' X

	// Concatenate strings using the "strings" package
	var strSlice2 = []string{"j","u","l","i","a","n"}
	var strBuilder strings.Builder
	for i := range strSlice2{
		strBuilder.WriteString(strSlice2[i])
	}
	var catStr2 = strBuilder.String()
	fmt.Printf("\n%v", catStr2)
	// String are unmutable in go, so once is created I cannot change it => catStr[0] = 'a' X
}