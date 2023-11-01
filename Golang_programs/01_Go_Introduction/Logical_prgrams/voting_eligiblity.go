// write a program to take input from user age and check whether check egibility for voting
package main

import "fmt"

func main() {
	var age int

	fmt.Print("Enter user age : ")
	fmt.Scan(&age)

	if age >= 18 {
		fmt.Println("You are eligible to vote")
	} else {
		fmt.Println("You are not eligible to vote")
	}
}
