package main

import "fmt"


func main() {
	//编译器自动判断数组长度
	arr3 := [...]int{2,3,4,5,6}
	fmt.Println(arr3)
	for _,v:= range arr3{
		fmt.Println(v)
	}
}
