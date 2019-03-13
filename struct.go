package main

import "fmt"


//定义一个矩形类型
type Rect struct {
	width, length float64
}

func main() {
	//定义一个Rect类型的变量rect
	var rect Rect

	rect.width = 100
	rect.length = 200

	fmt.Println(rect.width * rect.length)
	
}
