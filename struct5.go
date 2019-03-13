package main

import "fmt"

type Rect struct {
	width, length float64
}

//传入指针，然后乘2
func (rect *Rect) double_area() float64 {
	rect.width *= 2
	rect.length *= 2

	return rect.width * rect.length
}



func main() {
	var rect = new(Rect)
	rect.width = 100
	rect.length = 200
	//打印初始值
	fmt.Println(*rect)

	fmt.Println("Double Width:", rect.width, "Double Length:",rect.length, "DoubleArea:", rect.double_area())

	//再次打印
	fmt.Println(*rect)
}
