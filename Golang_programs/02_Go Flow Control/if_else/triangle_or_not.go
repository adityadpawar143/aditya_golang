// Determine the type of triangle based on angles
package main

import "fmt"

func main() {
	angle1, angle2, angle3 := 60, 60, 60

	if angle1+angle2+angle3 == 180 {
		if angle1 == angle2 && angle2 == angle3 {
			fmt.Println("The triangle is an equilateral triangle")
		} else if angle1 == angle2 || angle2 == angle3 || angle1 == angle3 {
			fmt.Println("The triangle is an isosceles triangle")
		} else {
			fmt.Println("The triangle is a scalene triangle")
		}
	} else {
		fmt.Println("Invalid triangle")
	}
}
