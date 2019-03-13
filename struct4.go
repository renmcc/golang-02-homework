package main

import "fmt"

type Rect struct {
	width, length float64
}

//矩形结构体方法 ,传入一个结构体指针变量rect ，第二个方法名，第三个返回值类型
func (rect *Rect) area() float64 {
	return rect.width * rect.length
}

func main() {
	//内置new方法创建一个结构体指针，结构体指针赋值不需要*
	var rect = new(Rect)
	rect.width = 100
	rect.length = 200

	fmt.Println("Width:", rect.width, "Length:", rect.length, "Area:", rect.area())
}
