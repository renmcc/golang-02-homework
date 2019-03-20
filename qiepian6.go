package main

import "fmt"

func main() {
	var arr []int
	for i := 0; i <= 100; i++ {
		arr = append(arr, i)
	}
	fmt.Println(arr)
}
