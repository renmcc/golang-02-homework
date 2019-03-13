package main

import "fmt"


//传入一个指针参数
func set_value(x_ptr *int) {
	*x_ptr = 100
}

func main() {
	//new方法申请一个存储整形的内存地址,赋值给x_ptr
	x_ptr := new(int)
	//赋值
	set_value(x_ptr)

	//打印x_ptr指向的地址
	fmt.Println(x_ptr)

	//打印x_ptr本身的地址
	fmt.Println(&x_ptr)

	//打印x_ptr指向地址的值
	fmt.Println(*x_ptr)

}
