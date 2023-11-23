package main

import "fmt"

func countDown(number int) {
	if number >= 0 {
		fmt.Println(number)
		countDown(number - 1)
	}
}

func main() {
	countDown(3)
}
