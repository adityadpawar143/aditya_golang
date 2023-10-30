// Write a program to take input from the user
package main

import "fmt"

func main() {
	var name string

	fmt.Print("Enter your name : ")

	fmt.Scan(&name)

	fmt.Printf("Hello %s", name)
}
