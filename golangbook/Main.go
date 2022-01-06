package main

import "fmt"

func main() {
	fmt.Println("Hello Olutoba")
	fmt.Println("Welcome to Go programming language!")
	fmt.Println(len("Olutoba"))
	fmt.Println("Hello World"[1])

	name := "Olutoba"
	age := 24
	fmt.Println(name, "is", age, "years old.")
	greetings := "Welcome back, Olutoba"
	fmt.Println(greetings)
	golang := "Go is an awesome programming language"
	fmt.Println(golang)

	wishlist := "I want to own a macbook pro and get a software engineering internship"
	fmt.Println(wishlist)

	const year = 2022
	fmt.Println(year)

	var (
		companies = "google, apple, twitter, meta, delivery hero"
		location  = "usa, canada, uk, germany"
	)
	fmt.Println(companies)
	fmt.Println(location)

	fmt.Print("Enter a number: ")
	var input float64
	fmt.Scanf("%f", &input)
	fmt.Println(input)

	fmt.Print("Enter fahrenheit value:")
	var fValue float64
	fmt.Scanf("%f", &fValue)
	fahrenheitToCelsius(fValue)

	feetToMetres(1)

	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	for i := 0; i < 6; i++ {
		switch i {
		case 0:
			fmt.Println("Zero")
		case 1:
			fmt.Println("One")
		case 2:
			fmt.Println("Two")
		case 3:
			fmt.Println("Three")
		case 4:
			fmt.Println("Four")
		case 5:
			fmt.Println("Five")
		default:
			fmt.Println("Unknown Number")
		}
	}

	for i := 1; i <= 100; i++ {
		if i%3 == 0 {
			fmt.Println("Fizz")
		} else if i%5 == 0 {
			fmt.Println("Buzz")
		} else if i%3 == 0 && i%5 == 0 {
			fmt.Println("FizzBuzz")
		}
	}

}

func fahrenheitToCelsius(fahrenheit float64) {
	var celsius float64
	celsius = (fahrenheit - 32) * 5 / 9
	fmt.Println(fahrenheit, "fahrenheit is", celsius, "celsius")
}

func feetToMetres(ft float64) {
	x := ft * 0.3048
	fmt.Println(ft, "feet is", x, "metres")
}
