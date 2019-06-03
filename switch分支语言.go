package main

import "fmt"

func main() {
	var w int
	fmt.Scan(&w)
	switch w {
	case 1:
		fmt.Println("星期一")
	case 2:
		fmt.Println("星期二")
	default:
		fmt.Println("错误输入")
	}
}