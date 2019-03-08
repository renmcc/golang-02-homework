package main

import "fmt"

func add(x,y int) int {
	z := x + y
	return z
}

func two(x, y int) (int, int) {
	return x +y, x * y
}

func main() {
	fmt.Println(add(1,3))
	fmt.Println(two(3,5))
}