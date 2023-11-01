// Check if a number is divisible by both 2 and 3
package main

import "fmt"

func main() {
	num := 12

	if num%2 == 0 && num%3 == 0 {
		fmt.Println("The number is divisible by both 2 and 3")
	} else {
		fmt.Println("The number is not divisible by both 2 and 3")
	}
}
