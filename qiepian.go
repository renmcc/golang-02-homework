package main

import "fmt"

//切片不用指定长度，是数组的视图可以直接修改
func printArray(arr []int) {
	arr[1] = 100

}

func main() {
	//编译器自动判断数组长度
	arr3 := [...]int{2,3,4,5,6}
	fmt.Println(arr3)
	printArray(arr3[:])
	fmt.Println(arr3)
}
