// Boolean Arrays
package main

import "fmt"

func main() {
	boolArray := []bool{true, false, true, true, false}
	for i, val := range boolArray {
		fmt.Printf("boolArray[%d]: %t\n", i, val)
	}
}
