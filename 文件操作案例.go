package main

import (
	"fmt"
	"io"
	"os"
)

/*
复制文件
 */

func main() {
	//文件拷贝 从一个文件拷贝到新的文件中
	fp1, err1 := os.Open("d:/EVE-20171007.iso")
	fp2, err2 := os.Create("d:/EVE-20171007_new.iso")

	if err1 != nil || err2 != nil {
		fmt.Println(err1,err2)
		return
	}
	defer fp1.Close()
	defer fp2.Close()

	//二进制文件只能按字节进行获取
	b := make([]byte, 1024)
	for {
		n,err := fp1.Read(b)
		if err == io.EOF {
			break
		}
		fp2.Write(b[:n])
	}
}