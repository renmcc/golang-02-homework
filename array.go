package main

import "fmt"

func main() {
	var x [5]int
	x[0] = 2
	x[1] = 3
	x[2] = 5
	x[3] = 3
	x[4] = 6

	var sum int
	for _, elem := range x {
		sum += elem
	}
	fmt.Println(sum)
}
