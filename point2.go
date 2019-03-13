package main

import "fmt"

//不用指针只能修改函数内变量x
func change(x int) {
	x = 200
}

//传入指针修改变量
func change2(x *int) {
	*x = 200
}



func main() {
	var x int = 100
	fmt.Println(x)
	change(x)
	fmt.Println(x)

	//调用指针修改
	change2(&x)
	fmt.Println(x)


}
