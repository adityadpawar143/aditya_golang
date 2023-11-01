// Write a program of ascending sort in go lang
package main

import (
	"fmt"
	"sort"
)

func main() {
	unsortedSlice := []int{5, 3, 2, 1, 4}

	sort.Ints(unsortedSlice)

	fmt.Println(unsortedSlice)
}
