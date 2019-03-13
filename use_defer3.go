package main

import "fmt"

func main() {
	//defer用于在main函数结束后执行一次，只能用于函数
	defer func() {
		//recover方法用于获取panic中定义的信息
		msg := recover()
		fmt.Println(msg)
	}()

	fmt.Println("I am walking and singging..")
	//触发异常
	panic("It starts to rain cats and dogs")

}
