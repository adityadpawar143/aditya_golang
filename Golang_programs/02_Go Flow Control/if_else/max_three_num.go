// Find the maximum of three numbers
package main

import "fmt"

func main() {
	a, b, c := 10, 15, 12
	max := a

	if b > max {
		max = b
	}
	if c > max {
		max = c
	}
	fmt.Println("The maximum is:", max)
}
