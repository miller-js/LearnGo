package main

import "fmt"

// Define a custom struct type
type Rectangle struct {
	width, height float64
}

// Area is a method with a receiver of type 'Rectangle'
func (r Rectangle) Area() float64 {
	return r.width * r.height
}

func main() {
	// Initialize the struct
	rect := Rectangle{width: 10, height: 5}

	// Call the method using dot notation
	fmt.Println("Area:", rect.Area()) // Output: Area: 50
}
