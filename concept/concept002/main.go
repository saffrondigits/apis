package main

import "fmt"

func main() {
	total := sum(1, 2, 3, 4, 5, 5, 12, 13, 14, 15, 16, 16)
	fmt.Println(total)
}

func sum(value ...int) int {
	sum := 0

	for i := 0; i < len(value); i++ {
		sum = sum + value[i]
	}

	return sum
}
