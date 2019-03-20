package main

import "fmt"

//需要传入指针
func printArray(arr *[5]int) {
	arr[1] = 100

}

func main() {
	//编译器自动判断数组长度
	arr3 := [...]int{2,3,4,5,6}
	fmt.Println(arr3)
	printArray(&arr3)
	fmt.Println(arr3)
}
