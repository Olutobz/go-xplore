package main

import "fmt"

/* Author: Onikoyi Damola Olutoba
   date: 06/01/2022
*/
func main() {
	var x [5]int
	x[4] = 24
	fmt.Println(x)
	findAverage()

	y := [3]float64{12, 14, 18}
	sum := 0
	for _, value := range y {
		fmt.Println(value)
		sum += int(value)
	}
	fmt.Println("sum is:", sum)

}

func findAverage() {
	var total float64 = 0
	var nums [5]float64
	nums[0] = 14
	nums[1] = 68
	nums[2] = 22
	nums[3] = 94
	nums[4] = 14
	for i := 0; i < len(nums); i++ {
		total += nums[i]
	}
	fmt.Println("average is:", total/float64(len(nums)))
}
