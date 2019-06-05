package main

import "fmt"

var sum1 int = 1

func test4(n int) {
	if n==1 {
		return
	}
	test4(n-1)
	sum1*=n
}

func main() {
	test4(100)
	fmt.Println(sum1)
}