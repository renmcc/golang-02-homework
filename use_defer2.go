package main

import (
	"fmt"
	"os"
	"bufio"
	"strings"
)


func main() {
	fname := "e:\\OpsManage\\pyvenv.cfg"
	f, err := os.Open(fname)
	//当主函数结束后关闭文件描述符
	defer f.Close()

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	bReader := bufio.NewReader(f)

	for {
		line, ok := bReader.ReadString('\n')
		if ok != nil {
			break
		}
		fmt.Println(strings.Trim(line, "\r\n"))
	}

}
