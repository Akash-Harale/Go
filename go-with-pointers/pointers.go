package main

import "fmt"

func main() {
	age := 32 //regular variable declaration

	agePointer := &age //pointer variable declaration , &age gives the memory address of age.

	fmt.Print(agePointer) // prints the memory address of age.
}
