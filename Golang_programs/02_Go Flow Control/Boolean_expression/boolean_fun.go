// Boolean Function
package main

import "fmt"

func isEven(num int) bool {
	return num%2 == 0
}

func main() {
	num := 4
	if isEven(num) {
		fmt.Println(num, "is even")
	} else {
		fmt.Println(num, "is odd")
	}
}
