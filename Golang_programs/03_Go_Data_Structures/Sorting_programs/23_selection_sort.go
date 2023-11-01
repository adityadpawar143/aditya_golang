// selection sort
package main

import (
	"fmt"
)

func selectionSort(array []int) []int {
	for i := 0; i < len(array)-1; i++ {
		minIndex := i
		for j := i + 1; j < len(array); j++ {
			if array[j] < array[minIndex] {
				minIndex = j
			}
		}
		// Swap the minimum element with the current element.
		array[i], array[minIndex] = array[minIndex], array[i]
	}
	return array
}

func main() {
	array := []int{5, 3, 2, 1, 4}

	sortedArray := selectionSort(array)

	fmt.Println(sortedArray)
}
