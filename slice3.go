package main

import "fmt"

func main() {
	//定义一个切片
	var arr1 = make([]int, 5, 10)

	for i := 0; i < len(arr1); i++ {
		arr1[i] = i
	}

	fmt.Println(arr1)

	//内置方法append追加元素到arr1
	arr1 = append(arr1, 5)

	//打印切片长度和内容
	fmt.Println("Capacity:", cap(arr1), "Length:", len(arr1))
	fmt.Println(arr1)
}
