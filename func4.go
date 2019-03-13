package main

import "fmt"

//闭包函数,返回偶数
func createEvenGenerator() func() uint {
	i := uint(0)
	return func() uint {
		var retVal uint
		retVal = i
		i += 2
		return retVal
	}
}

func main() {
	nextEven := createEvenGenerator()	
	//循环50次去除100以内的偶数
	for i := 0; i <= 50; i++ {
		fmt.Println(nextEven())
	}
}

