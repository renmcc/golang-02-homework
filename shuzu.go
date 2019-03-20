package main

import "fmt"

//必须指明数组的长度
func printArray(arr [5]int) {
	for i,v := range arr {
		fmt.Println(i,v)
	}
}

func main() {
	//编译器自动判断数组长度
	arr3 := [...]int{2,3,4,5,6}
}
