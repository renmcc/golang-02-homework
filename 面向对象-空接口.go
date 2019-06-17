package main

import "fmt"

func main111055() {
	//空接口 可以接收任意类型数据
	var i interface{}

	i = 10
	fmt.Println(i)
	fmt.Printf("%p\n", &i)

	i = "hello world"
	fmt.Println(i)
	fmt.Printf("%p\n", &i)

	arr := [3]int{1,2,3}
	i = arr
	fmt.Println(i)
	fmt.Printf("%p\n", &i)
}

func main() {
	//空接口切片,存储任意类型
	var i []interface{}
	i = append(i,1,2,3,4,5,"hello world",[]int{1,2,3,4,5})
	for _,v := range i{
		fmt.Println(v)
	}
}