package main

import "fmt"

type Rect struct {
	width, length float64
}

//矩形结构体方法 ,传入一个结构体变量rect ，第二个方法名，第三个返回值类型
func (rect Rect) area() float64 {
	return rect.width * rect.length
}

func main() {
	var rect = Rect{100, 222}

	fmt.Println("Width:", rect.width, "Length:", rect.length, "Area:", rect.area())
}
