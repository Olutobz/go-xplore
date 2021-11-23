package main

/*Author: Onikoyi Damola Olutoba
  date: 16/11/2021
*/
import "fmt"

var j = 16

/* We have 3 ways of declaring variables in go
1. var i int
2. var i = 20
3. i := 20
*/

// multiple variable declarations
var (
	actorName   string = "Dwayne Johnson"
	idNumber    int    = 44
	actressName string = "Jessica Alba"
	season      int    = 6
)

var (
	name = "Onikoyi Damola Olutoba"
	age  = 24
)

var googleHttpURL string = "https://www.google.com"

func main() {
	var i = 40
	fmt.Println(i)
	fmt.Printf("%v, %T", j, j)
	println()
	fmt.Println(actorName, idNumber, actressName, season)
	fmt.Println(name, age)
	fmt.Println("Url: " + googleHttpURL)

	var j float32
	j = float32(i)
	fmt.Printf("%v, %T", j, j)
}
