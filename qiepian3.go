package main

import "fmt"

func main() {
	//创建一个数组
	a := [3]int{1,2,3}
	//创建一个切片b
	var b []int = a[:]
	//修改切片值
	for i := range b {
		b[i]++
	}
	//再次查看数组的值
	fmt.Println(a)
}
