package main

import (
	"fmt"
	"os"
	"bufio"
)

func printFile(filename string) {
	if file, err := os.Open(filename); err != nil {
		panic(err)
	}else {
		//使用bufio按行读取文件
		scanner := bufio.NewScanner(file)

		for scanner.Scan() {
			fmt.Println(scanner.Text())
		}
	}
}

func main() {
	printFile("abc.txt")
}
