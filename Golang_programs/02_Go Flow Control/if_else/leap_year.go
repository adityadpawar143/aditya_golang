// Determine if a year is a leap year
package main

import "fmt"

func main() {
	year := 2024

	if (year%4 == 0 && year%100 != 0) || (year%400 == 0) {
		fmt.Println(year, "is a leap year")
	} else {
		fmt.Println(year, "is not a leap year")
	}
}
