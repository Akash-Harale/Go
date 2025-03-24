package main

import "fmt"

func main() {
	age := 32 //regular variable declaration

	agePointer := &age //pointer variable declaration , &age gives the memory address of age.

	fmt.Print("age: ", agePointer) // ths pointer value is not useful here, but it can be used in other places. such as in function calls.

	fmt.Print("age: ", *agePointer) //by adding * before the pointer, we can access the value of the variable it points to not the memory address. this is called dereferencing.

}
