package main

import "fmt"

func main() {
	fmt.Println("Hello Olutoba")
	fmt.Println("Welcome to Go programming language!")
	fmt.Println(len("Olutoba"))
	fmt.Println("Hello World"[1])

	hello := "Hello, Olutoba"
	fmt.Println(hello)
	golang := "Go is an awesome programming language"
	fmt.Println(golang)

	laptopName := "I want to own a macbook pro"
	fmt.Println(laptopName)

	const year = 2022
	fmt.Println(year)

	var (
		companies = "google, apple, twitter, meta, delivery hero"
		location = "usa, canada, uk, germany"
	)
	fmt.Println(companies)
	fmt.Println(location)

}
