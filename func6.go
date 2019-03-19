package main

import (
	"fmt"
	"reflect"
	"runtime"
	"math"
)

//函数式编程
func apply(op func(int,int) int, a, b int) int {
	//反射
	p := reflect.ValueOf(op).Pointer()
	opName := runtime.FuncForPC(p).Name()
	fmt.Printf("Calling function %s with args" + "(%d,%d)", opName, a, b )
	return op(a, b)
}

//重写match.Pow()
func pow(a,b int) int {
	return int(math.Pow(float64(a), float64(b)))
}

func main() {
	//fmt.Println(apply(pow, 2, 4))
	//匿名函数
	fmt.Println(apply(
		func(a int, b int) int{
			return int(math.Pow(
				float64(a), float64(b)))
		}, 2, 4))

}
