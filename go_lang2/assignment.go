// use of assignment operators in golang
package main

import "fmt"

func main() {
	var a int = 45
	var b int = 50

	// “=”(Simple Assignment)
	a = b
	fmt.Println(a)

	// “+=”(Add Assignment)
	a += b
	fmt.Println(a)

	//“-=”(Subtract Assignment)
	a -= b
	fmt.Println(a)

	// “*=”(Multiply Assignment)
	a *= b
	fmt.Println(a)

	// “/=”(Division Assignment)
	a /= b
	fmt.Println(a)

	// “%=”(Modulus Assignment)
	a %= b
	fmt.Println(a)

}
