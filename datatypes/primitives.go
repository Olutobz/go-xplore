package main

import (
	"fmt"
)

func main() {

	// Boolean type, default is false (0)
	var name bool
	fmt.Printf("%v, %T\n", name, name)
	b1 := true
	b2 := false
	fmt.Printf("%v, %T\n", b1, b1)
	fmt.Printf("%v, %T\n", b2, b2)

	// Numeric types: int8, int16, int32, int64
	n := 42
	fmt.Printf("%v, %T", n, n)

	// Unsigned integers: uint8, uint16, uint32
	var n2 uint16 = 42
	fmt.Printf("%v, %T\n", n2, n2)

}
