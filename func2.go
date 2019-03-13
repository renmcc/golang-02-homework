package main

import "fmt"

//可变长参数，函数以数组传入，可迭代,必须是最后一个参数
func sum(arr ...int) int {
	sum := 0

	for _, val := range arr {
		sum += val
	}
	return sum
}

func main() {
	fmt.Println(sum(1))	
	fmt.Println(sum(1,2,3,4,5))	
}
