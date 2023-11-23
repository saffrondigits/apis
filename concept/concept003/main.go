package main

import "fmt"

// outer function
func greet() func() string {
	// variable defined outside the inner function
	name := "John"

	// return a nested anonymous function
	return func() string {
		name = "Hi " + name
		return name
	}
}

func main() {
	// call the outer function
	message := greet()

	greet := message()
	// call the inner function
	fmt.Println(greet)
}
