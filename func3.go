package main

import "fmt"

func sum(base int, arr ...int) int {
	sum := base

	for _, val := range arr {
		sum += val
	}
	return sum
}

func main() {
	var arr1 = []int{1,2,3,4,5}
	//把切片解包后传入函数sum
	fmt.Println(sum(200, arr1...))
}
