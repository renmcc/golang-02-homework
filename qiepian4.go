package main

import "fmt"

func main() {
	//创建数组
	arr := [...]string{"zero","one"}
	//创建切片数组
	a := []string{"aa","bb","cc"}
	b := []string{"1","2"}
	//添加一个元素
	a = append(a, "dd")
	//使用...将一个切片追加到另一个切片
	a = append(a, b...)
	a = append(a, arr[:]...)
	fmt.Println(a)
}
