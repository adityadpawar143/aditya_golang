// Write a program of descending sort in go lang
package main

import (
	"fmt"
	"sort"
)

func main() {
	unsortedSlice := []int{5, 3, 2, 1, 4}

	sort.Sort(sort.Reverse(sort.IntSlice(unsortedSlice)))

	fmt.Println(unsortedSlice)
}
