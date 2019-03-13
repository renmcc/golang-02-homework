package main

import "fmt"

func main() {
	var x = [...]int{1,1,2,3}

	var sum int
	for _, elem := range x {
		sum += elem
	}
	fmt.Println(sum)
}
