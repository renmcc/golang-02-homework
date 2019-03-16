package main

import (
	"flag"
	"fmt"
)

func main() {
	filename := flag.String("file", "", "输入文件名")
	flag.Parse()
	fmt.Println("filename", *filename)
	//多余的参数保存到列表里
	fmt.Println("args", flag.Args())
}
