// Boolean Pointers
package main

import "fmt"

func main() {
	b := true
	var p *bool
	p = &b

	fmt.Println("Value:", b)
	fmt.Println("Pointer value:", *p)
}
