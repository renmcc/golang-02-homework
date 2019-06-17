package main

import "fmt"

func test1(args ...int) {
	fmt.Println(args)

}

func test(args ...int) {
	test1(args...)

}

func main() {
	test(1,2,3,4)
}

