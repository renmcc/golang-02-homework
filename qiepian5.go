package main

import "fmt"

func main() {
	//创建数组
	arr := [...]string{"zero","one", "two", "three", "four", "fire", "six", "sreen"}
	//当我们只用sel一个变量时，arr也会在内存中
	sel := arr[1:4]
	//创建一个空的切片
	sel2 := make([]string, len(sel))
	//复制sel到sel2，arr会被垃圾回收
	copy(sel2, sel)
	//查看容量
	fmt.Println(cap(sel))
	fmt.Println(cap(sel2))
	fmt.Println(sel2)
}
