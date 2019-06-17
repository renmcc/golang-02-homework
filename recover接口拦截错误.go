package main

import "fmt"

func demo(i int) {
	var arr [10]int

	//通过匿名函数和recover()进行错误拦截
	//recover可以从panic异常中重新获取控制权
	//返回值为接口类型
	defer func(){
		err := recover()
		if err != nil {
			fmt.Println(err)
		}
	}()

	//如果传递值超过数组下标，导致数组下标越界panic异常
	arr[i] = 100
	fmt.Println(arr)
}

func demo1() {
	fmt.Println("hello world")
}

func main() {
	demo(11)
	demo1()
}