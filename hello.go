package main

/*Author: Onikoyi Damola Olutoba
  date: 16/11/2021
*/
import "fmt"

var j = 16


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

var googleHttpUrl string = "https://www.google.com"

func main() {
	var i = 40
	fmt.Println(i)
	fmt.Printf("%v, %T", j, j)
	println()
	fmt.Println(actorName, idNumber, actressName, season)
	fmt.Println(name, age)
	fmt.Println("Url: " + googleHttpUrl)

	var j float32
	j = float32(i)
	fmt.Printf("%v, %T", j, j)
}