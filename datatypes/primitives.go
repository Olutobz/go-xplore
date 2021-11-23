package main

import "fmt"

func main() {

	// Boolean type, values are T or F, default is false (0)
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
	op1 := 10               // 1 0 1 0
	op2 := 3                // 0 0 1 1
	fmt.Println(op1 & op2)  // 0 0 1 0
	fmt.Println(op1 | op2)  // 1 0 1 1
	fmt.Println(op1 ^ op2)  // 1 0 0 1
	fmt.Println(op1 &^ op2) // 0 1 0 0

	// Bit shifting
	d := 8              // 2 ^ 3
	fmt.Println(d << 3) // 8 * 2 ^ 3 = 64
	fmt.Println(d >> 3) // 8 / 2 ^ 3 = 2 ^ 0 -> 1

	// Floating-point numbers
	f := 3.14
	f = 13.7e72
	f = 2.1e14
	fmt.Printf("%v, %T\n", f, f)

	// Complex numbers
	var num complex64 = 1 + 2i
	fmt.Printf("%v, %T\n", num, num)
	fmt.Printf("%v, %T\n", real(num), real(num))
	fmt.Printf("%v, %T\n", imag(num), imag(num))

	str := "My name is olutoba"
	fmt.Printf("%v,  %T\n", str, str)
	fmt.Printf("%v,  %T\n", string(str[6]), str[2])

	dogName := "max"
	fmt.Println("My dog's name is ", dogName)

	fmt.Print("Enter a number: ")
	var input float64
	fmt.Scanf("%f", &input)
	output := input * 2
	fmt.Println(output)

	// control flow
	for i := 0; i < 10; i++ {
		fmt.Print(i)
		fmt.Println()
	}

	for i := 1; i <= 20; i++ {
		if i%2 == 0 {
			fmt.Println(i, "even")
		} else {
			fmt.Println(i, "odd")
		}
	}

	// Arrays
	var x [5]int
	x[0] = 16
	x[1] = 2
	x[2] = 4
	x[3] = 100
	total := 0
	for i := 0; i < len(x); i++ {
		total += x[i]
	}
	fmt.Println("total number is: ", total)

	var floatArray = [10]float64{17, 42, 35, 98}
	fmt.Println(floatArray)

	// Slices -> A segment of an array
	slice := make([]float64, 5)
	fmt.Println(slice)

	// Maps (Key -> Value pair)
	y := make(map[int]int)
	y[1] = 24
	fmt.Println(x[1])

	elements := make(map[string]string)
	elements["H"] = "Hydrogen"
	elements["He"] = "Helium"
	elements["Li"] = "Lithium"
	elements["Be"] = "Beryllium"
	elements["B"] = "Boron"
	elements["C"] = "Carbon"
	elements["N"] = "Nitrogen"
	elements["O"] = "Oxygen"
	elements["F"] = "Fluorine"
	elements["Ne"] = "Neon"
	fmt.Println(elements)

	fmt.Println(add(1, 2, 3))

	fmt.Println("factorial of 5! is: ",factorial(5))
}

// functions
func average(x int) {
	fmt.Println("average method")
}

func add(args ...int) int {
	total := 0
	for _, arg := range args {
		total += arg
	}
	return total
}

func factorial(x uint) uint {
	if x == 0 {
		return 1
	}
	return x * (factorial(x - 1))
}
