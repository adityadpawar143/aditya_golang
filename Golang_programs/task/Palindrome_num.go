// write a program to check number is palindrome or not
package main

import (
	"fmt"
)

func isPalindrome(num int) bool { /* This line declares a function called isPalindrome(),
	                                  which takes a number as input and returns a boolean value
									  indicating whether or not the number is a palindrome.*/

	reversedNum := 0
	temp := num
	for temp > 0 { // The for loop will iterate until the temp variable is greater than zero.

		reversedNum = reversedNum*10 + temp%10 //It does this by multiplying the reversedNum variable by 10 and
		// then adding the remainder of the temp variable divided by 10.

		temp /= 10 //  This line divides the temp variable by 10. This removes the most significant digit from the number.
	}
	return num == reversedNum
}

func main() {
	var num int
	fmt.Scanln(&num)

	if isPalindrome(num) {
		fmt.Println("The number is a palindrome")
	} else {
		fmt.Println("The number is not a palindrome")
	}
}
