package main
import (
	"fmt"
	
)

// // Structs and interfaces
// type gasEngine struct {
// 	mpg uint8
// 	gallons uint8
// 	// ownerInfo owner 
// 	owner // Here I can use the owner directly so I don't need to access the ownerInfo but direcly the owner.name
// }

// type owner struct {
// 	name string
// }

// func main() {
// 	// var myEngine gasEngine = gasEngine{mpg: 25, gallons: 15}
// 	 var myEngine gasEngine = gasEngine{25, 15, owner{"Alex"}} // I can omit the names
// 	// I can assign them like this: myEngine.mpg = 20
// 	//fmt.Println(myEngine.mpg, myEngine.gallons, myEngine.ownerInfo.name)
// 	fmt.Println(myEngine.mpg, myEngine.gallons, myEngine.owner.name)
// }


// // Structs and interfaces
// type gasEngine struct {
// 	mpg uint8
// 	gallons uint8
// }

// func main() {
// 	var myEngine gasEngine = gasEngine{25, 15} 
// 	// var myEngine2 = struct{
// 	// 	mpg uint8
// 	// 	gallons uint8
// 	// }{25,15}            This is not reusable, if want to create another struct I would have to re-write it
// 	// var myEngine3 = struct{
// 	// 	mpg uint8
// 	// 	gallons uint8
// 	// }{25,15} 
// 	fmt.Println(myEngine.mpg, myEngine.gallons)
// }



// Structs and interfaces
// Methods
type gasEngine struct {
	mpg uint8
	gallons uint8
}

type electricEngine struct {
	mpkwh uint8
	kwh uint8
}

func (e gasEngine) milesLeft() uint8 { // The (e gasEngine) is to assing this function to the gasEngine type
	return e.gallons*e.mpg
}

func (e electricEngine) milesLeft() uint8{
	return e.kwh*e.mpkwh
}

// Make an interface
type engine interface {
	milesLeft() uint8 
}

func canMakeIt(e engine, miles uint8){
	if miles <= e.milesLeft(){
		fmt.Println("You can make it there!")
	} else {
		fmt.Println("Need to fuel up first!")
	}
}

func main() {
	var myEngine gasEngine = gasEngine{25, 15} 
	canMakeIt(myEngine, 50)

	var myOtherEngine electricEngine = electricEngine{25,15}
	canMakeIt(myOtherEngine, 50)

	// This can now take any type of engine, this is beautiful
}