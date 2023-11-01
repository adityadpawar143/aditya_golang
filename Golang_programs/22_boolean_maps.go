// Boolean Maps
package main

import "fmt"

func main() {
	boolMap := map[string]bool{
		"A": true,
		"B": false,
		"C": true,
	}

	for key, val := range boolMap {
		fmt.Printf("%s: %t\n", key, val)
	}
}
