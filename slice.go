package main

import "fmt"

func main() {
	//定义切片x和y
	var x = make([]float64, 5)
	var y = make([]float64, 5, 10)
	//打印切片的容量和长度
	fmt.Println("Capcity:", cap(x), "Length", len(x))
	fmt.Println("Capcity:", cap(y), "Length", len(y))
        //给切片赋值
	for i := 0; i < len(x); i++ {
		x[i] = float64(i)
	}
	fmt.Println(x)

	for i := 0; i < len(y); i++ {
		y[i] = float64(i)
	}
	fmt.Println(y)
}
