package main

import "fmt"

func jiecheng(sub int) int {
	if sub==0{
		return 1
	}
	return sub * jiecheng(sub -1)
}

func main() {
	for i:=0;i<=7;i++{
		value := jiecheng(i)
		fmt.Printf("%d的阶乘是：%d\n", i, value)
	}
}
