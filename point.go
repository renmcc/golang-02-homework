package main

import "fmt"

func main() {
	var x int
	var x_ptr *int

	x = 10
	//&取一个变量的地址
	x_ptr = &x


	fmt.Println(x)
	fmt.Println(x_ptr)
	//*取一个指针变量所指向的地址的值
	fmt.Println(*x_ptr)
	//获取指针的地址
	fmt.Println(&x_ptr)

}
