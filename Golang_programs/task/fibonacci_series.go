// write program to find fibonacci series upto specific number
package main

import "fmt"

func fibonacci(n int) {
	var a, b = 0, 1
	for i := 0; i < n; i++ {
		fmt.Print(a, " ")
		a, b = b, a+b
	}
}

func main() {
	fmt.Println("Fibonacci Series:")
	fibonacci(10)
}
