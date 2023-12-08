// Sure! Let's imagine you have a magical box that can change its size to fit exactly 
// whatever type of toy you want to put inside it. If you have a small toy car, the box
// shrinks to hold just the car. If you have a big teddy bear, the box expands to snugly fit
// the teddy bear. This box is special because it can hold any one toy, no matter its size, 
// and it always fits just right.
// In Go, generics work like this magical box. Generics allow you to create a function, or a 
// type, that can work with any type of data. Here's how you might explain it with code:




package main

import "fmt"

// This is our magical box. It's a generic type, which means it can become a box for any type of toy.
// 'ToyType' is a placeholder for the actual toy you'll put in the box later.
type MagicalBox[ToyType any] struct {
    // The content of the box can be any type, just like the box can hold any toy.
    Content ToyType
}

// This function allows you to put a toy into the magical box.
// 'ToyType' here means it can be used with any toy, not just a specific type.
func PutToyInBox[ToyType any](toy ToyType) MagicalBox[ToyType] {
    // We create a new magical box for your toy.
    return MagicalBox[ToyType]{Content: toy}
}

// This function lets you see what's inside your magical box.
func (mb MagicalBox[ToyType]) RevealToy() ToyType {
    // It shows you the toy in the box.
    return mb.Content
}

func main() {
    // Let's put a toy car in our magical box. The box becomes a box for toy cars.
    carBox := PutToyInBox("toy car")
    fmt.Println("The box contains a:", carBox.RevealToy())

    // Now, let's put a teddy bear in the box. The box changes to fit a teddy bear.
    bearBox := PutToyInBox(123) // This time, our toy is a number, just to show it can be anything!
    fmt.Println("The box contains a:", bearBox.RevealToy())
}
