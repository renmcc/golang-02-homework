package main

import "fmt"

func main() {
	var arr []int
	for i := 0; i <= 100; i++ {
		arr = append(arr, i)
	}
	//删掉11-20的索引
	arr = append(arr[:11], arr[21:101]...)
	fmt.Println(arr)
}
