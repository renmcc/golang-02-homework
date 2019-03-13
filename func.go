package main

import "fmt"

func slice_sum(arr []int) (int, float64) {
	sum := 0
	avg := 0.0

	//求和
	for _, elem := range arr {
		sum += elem
	}
	//求平均值
	avg = float64(sum) / float64(len(arr))

	return sum, avg
}

func main() {
	var arr1 = []int{1,2,3,4,5}
	var arr2 = []int{1,2,3,4,5,5,6,4}

	fmt.Println(slice_sum(arr1))
	fmt.Println(slice_sum(arr2))
}
