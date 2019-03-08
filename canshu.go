package main

import (
	"fmt"
	"os"
	"strconv"
)

func print_args() string {
	var s ,sep string
	for i :=1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	return s
}

func addarg() int {
	arg1, _ := strconv.Atoi(os.Args[1])
	arg2, _ := strconv.Atoi(os.Args[2])
	return arg1 + arg2
}

func main() {
	fmt.Println(print_args())
	fmt.Println(addarg())
}
