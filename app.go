package main

import "fmt"

func main() {
	age := 32

	var agePointer *int // pointer variable declaration *int is used to declare a pointer to an integer.

	agePointer = &age // assigning the memory address of age to agePointer.

	fmt.Print("age: ", *agePointer, "\n") // dereferencing. This prints the value of the variable age that agePointer points to.

	*agePointer = 33 // changing the value of the variable that agePointer points to.

	// getAdultAgeYear(age) --> This won't work because age is a pointer.
	// getAdultAgeYear(agePointer) --> This will work because agePointer is a pointer to an integer.

	// OR

	adultYear := getAdultAgeYear(&age) // passing the memory address of age to the function.

	fmt.Print("adultYear: ", adultYear, "\n ")
}

func getAdultAgeYear(age *int) int { // you can't perform arithmetic on pointers as we do with regular variables.

	// return age-18  --> we can't do this here as age is a pointer.
	return *age - 18 // dereferencing to get the actual value of age.
}
