package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	//最好的方法
	fmt.Println(strings.Join(os.Args[1:], ""))
	//第二种方法
	ret := ""
	for _,arg := range os.Args[1:] {
		ret += arg
	}
	fmt.Println(ret)
	//第三种方法
	ret2 := ""
	for i := 1; i < len(os.Args); i++ {
		ret2 += os.Args[i]
	}
	fmt.Println(ret2)
}
