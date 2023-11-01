// if else
package main

import "fmt"

func main() {
	a, b := true, false

	if a && b {
		fmt.Println("Both a and b are true")
	} else if a || b {
		fmt.Println("At least one of a or b is true")
	} else {
		fmt.Println("Both a and b are false")
	}
}
