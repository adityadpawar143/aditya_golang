// write a program to find given number is Armstrong number
package main

import (
	"fmt"
	"math"
)

func isArmstrongNumber(number int) bool {
	numDigits := int(math.Log10(float64(number))) + 1
	temp := number
	sum := 0

	for temp != 0 {
		digit := temp % 10
		sum += int(math.Pow(float64(digit), float64(numDigits)))
		temp /= 10
	}

	return sum == number
}

func main() {
	number := 153 // Change this to the desired number to check if it's an Armstrong number
	if isArmstrongNumber(number) {
		fmt.Printf("%d is an Armstrong number\n", number)
	} else {
		fmt.Printf("%d is not an Armstrong number\n", number)
	}
}
