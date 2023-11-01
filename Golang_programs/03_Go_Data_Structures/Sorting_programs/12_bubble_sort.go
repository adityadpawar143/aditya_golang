// Write bubble sort program in go lang
package main

import "fmt"

func bubbleSort(slice []int) {
	swapped := true

	for swapped {
		swapped = false

		for i := 0; i < len(slice)-1; i++ {
			if slice[i] > slice[i+1] {
				slice[i], slice[i+1] = slice[i+1], slice[i]
				swapped = true
			}
		}
	}
}

func main() {
	slice := []int{5, 3, 1, 2, 4}
	bubbleSort(slice)
	fmt.Println(slice)
}
