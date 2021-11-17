package main

import "fmt"

func main() {

	// Boolean type, valuees are T or F, default is false (0)
	var name bool
	fmt.Printf("%v, %T\n", name, name)
	b1 := true
	b2 := false
	fmt.Printf("%v, %T\n", b1, b1)
	fmt.Printf("%v, %T\n", b2, b2)

	// Numeric types: int8, int16, int32, int64
	n := 42
	fmt.Printf("%v, %T\n", n, n)

	// basic arithmetic
	a := 10
	b := 3
	fmt.Println("addition:", a+b)
	fmt.Println("subtract:", a-b)
	fmt.Println("multiply", a*b)
	fmt.Println("divide:", a/b)
	fmt.Println("modulo:", a%b)

	// Unsigned integers: uint8, uint16, uint32
	var n2 uint16 = 42
	fmt.Printf("%v, %T\n", n2, n2)

	// Bit operators
	op1 := 10               //1 0 1 0
	op2 := 3                //0 0 1 1
	fmt.Println(op1 & op2)  // 0 0 1 0
	fmt.Println(op1 | op2)  // 1 0 1 1
	fmt.Println(op1 ^ op2)  // 1 0 0 1
	fmt.Println(op1 &^ op2) // 0 1 0 0

	// Bit shifting
	d := 8              // 2 ^ 3
	fmt.Println(d << 3) // 8 * 2 ^ 3 = 64
	fmt.Println(d >> 3) // 8 / 2 ^ 3 = 2 ^ 0 -> 1

	// Floating-point numbers
	f1 := 3.14
	f1 = 13.7e72
	f1 = 2.1e14
	fmt.Printf("%v, %T\n", f1, f1)

	// Complex numbers
	var num complex64 = 1 + 2i
	fmt.Printf("%v, %T\n", num, num)
	fmt.Printf("%v, %T\n", real(num), real(num))
	fmt.Printf("%v, %T\n", imag(num), imag(num))

	str := "My name is olutoba"
	fmt.Printf("%v,  %T\n", str, str)
	fmt.Printf("%v,  %T\n", string(str[6]), str[2])

}
