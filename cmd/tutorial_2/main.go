package main
import "fmt"
import "unicode/utf8"

func main() {
	var intNum int = 32767
	intNum = intNum + 1
	fmt.Println(intNum)

	var floatNum float64 = 12345678.9
	fmt.Println(floatNum)

	var floatNum32 float32 = 10.1
	var intNum32 int32 = 2
	var result float32 = floatNum32 + float32(intNum32)
	fmt.Println(result)

	var intNum1 int = 3
	var intNum2 int = 2
	fmt.Println(intNum1/intNum2)
	fmt.Println(intNum1%intNum2)

	var myString string = "Hello, world!"
	fmt.Println(myString)

	var myString2 string = "Hello, \nworld!"
	fmt.Println(myString2)

	var myString3 string = `Hello
	World!`
	fmt.Println(myString3)

	var myString4 string = "Hello" + " " + "World!"
	fmt.Println(myString4)

	fmt.Println(len("test")) //This does not print the lengs of the string but the number of bites 10110111

	//To get the lengh of a string and the number of characters import "unicode/utf8"
	fmt.Println(utf8.RuneCountInString("Y"))

	var myRune rune = 'a'
	fmt.Println(myRune) //In Go, when you assign a character literal to a rune variable, it represents the Unicode code point of that character. The character 'a' corresponds to the Unicode code point 97 in the ASCII character encoding.

	var myBoolean bool = false
	fmt.Println(myBoolean)

	var intNum3 int
	fmt.Println(intNum3) 
	//the default value for (uint, uint8, uint16, uint32, uint64), signed ints (int, int8, int16, int32, int64), floating-point numbers (float32, float64), and rune for Unicode code points is = 0
	//the default value for strings is = ''
	//the default value for bool is = false

	//var myVar = "text" //in this way the type is inferred
	myVar := "text" //This is how it can be simplified
	fmt.Println(myVar)

	//var var1, var2 int = 1,2 //in this way the type is inferred
	var1, var2 := 1, 2 //This is how it can be simplified
	fmt.Println(var1, var2)

	//Add the type whhen it is not obvious
	// var myVar2 string = foo()
	// fmt.Println(myVar2)

	//cannot assign to myConst (constant "Const Value" of type string)
	// const myConst string = "Const Value"
	// myConst = "Another const value"

	const myConst string = "const value" //Const variables have to be initialized
	fmt.Println(myConst)

	const pi float32 = 3.1415
	fmt.Println(pi)
}

