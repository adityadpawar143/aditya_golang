// write a program to take input from the user and  find prime numbers in given range
package main

import (
	"fmt"
	"math"
)

func isPrime(n int) bool {
	if n <= 1 {
		return false
	}
	for i := 2; i <= int(math.Sqrt(float64(n))); i++ { //The for loop iterates over all the numbers from 2 to the square root of the given number.
		if n%i == 0 {
			return false
		}
	}
	return true
}

func printPrimeNumbers(start, end int) {
	for i := start; i <= end; i++ {
		if isPrime(i) {
			fmt.Printf("%d ", i)
		}
	}
	fmt.Println()
}

func main() {
	var start, end int
	fmt.Println("Enter the start and end of the range:")
	fmt.Scanln(&start, &end)

	printPrimeNumbers(start, end)
}
