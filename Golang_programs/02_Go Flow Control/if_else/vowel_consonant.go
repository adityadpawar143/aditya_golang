// Check if a character is a vowel or a consonant
package main

import "fmt"

func main() {
	char := 'a'

	if char == 'a' || char == 'e' || char == 'i' || char == 'o' || char == 'u' {
		fmt.Println("The character is a vowel")
	} else {
		fmt.Println("The character is a consonant")
	}
}
