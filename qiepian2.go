package main

import "fmt"


func main() {
	//创建一个数组
	a := [3]int{1,2,3}
	//创建一个切片b
	var b []int = a[0:2]
	//创建切片并且赋值
	c := []int{6,7,8}
	//用make创建一个切片,长度5，容量5
	d := make([]int, 5, 5)
	fmt.Printf("%T, %v", b,b)
	fmt.Println(c)
	fmt.Println(d)
}
