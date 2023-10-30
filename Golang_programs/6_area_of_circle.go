// write a program to take radius from user & print area of circle
package main

import "fmt"

func main() {
	var r float64
	fmt.Println("Enter the Radius: ")
	fmt.Scan(&r)

	const PI = 3.14

	area := PI * r * r

	fmt.Println("Radius is : ", r)
	fmt.Println("Area of Circle is :", area)
}
