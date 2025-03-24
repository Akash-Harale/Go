package main

import "fmt"

func main() {
	age := 32

	fmt.Print("Age: ", age, "\n")
	fmt.Print(getAdultAge(age), "\n")
}

func getAdultAge(age int) int {
	return age - 18
}
