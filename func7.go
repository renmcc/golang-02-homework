package main

import "fmt"

//可变参数
func sum(numbers ...int) int {
	s := 0
	for i, elem := range numbers {
		s += numbers[i]
		fmt.Println(elem)
	}
	return s
}

func main() {
	fmt.Println(sum(1,2,3,4,5,6))	
}
