package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	const filename = "abc.txt"
	//读取文件
	if contents, err := ioutil.ReadFile(filename); err != nil {
		fmt.Println(err)
	}else {
		fmt.Printf("%s\n", contents)
	}
}
