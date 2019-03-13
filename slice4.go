package main

import "fmt"

func main() {
	//定义一个切片，并赋值
	slice1 := []int{1,2,3,4,5,6}
	//定义一个空的切片
	slice2 := make([]int, 5, 10)

	//复制slice1元素到slice2
	copy(slice2, slice1)

	//因为slince2长度为5，所以只能添加前5个元素
	fmt.Println(slice2)
}
