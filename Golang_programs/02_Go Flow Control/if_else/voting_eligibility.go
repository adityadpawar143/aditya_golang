// Check eligibility for voting
package main

import "fmt"

func main() {
	age := 18

	if age >= 18 {
		fmt.Println("You are eligible to vote")
	} else {
		fmt.Println("You are not eligible to vote")
	}
}
