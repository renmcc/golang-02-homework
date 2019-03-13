package main

import "fmt"

func fibonacci(n int) int {
	retVal := 0

	if n == 1 {
		retVal = 1
	}else if n == 2 {
		retVal = 2
	}else {
		retVal = fibonacci(n-2) + fibonacci(n-1)
	}
	return retVal
}

func main(){
	for i := 1; i <= 100; i++ {
		fmt.Println(fibonacci(i))
	}
}
