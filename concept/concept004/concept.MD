In Go (Golang), recursion is a programming concept where a function calls itself. This technique is used to solve problems that can be broken down into smaller, repetitive problems.

Here’s an example of a simple recursive function in Go:


```
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
```


In this example, the countDown function calls itself, decrementing the number each time until number is no longer greater than or equal to 01. This is the base case that stops the recursion2.

Recursion can be of different types2:

Direct Recursion: The function calls itself directly2.
Indirect Recursion: The function calls another function, and this function, in turn, calls the original function2.
Tail Recursion: The recursive call is the last thing executed by the function2.
Head Recursion: The recursive call is the first statement in the function2.
Remember, it’s important to have a base case in a recursive function to prevent infinite recursion1. Without a base case, the function would keep calling itself indefinitely, which could lead to a stack overflow error3.