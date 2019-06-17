package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	//os.Create如果文件不存在创建，如果有会覆盖原文件
	//fp, err := os.Create("D:/a.txt")
	//以只读方式打开文件
	//fp, err := os.Open("D:/a.txt")

	//三个参数：1、打开的文件
	//2、打开模式，os.O_WRONLY只写 os.O_RDONLY只读 os.O_RDWR可读可写 os.O_APPEND追加模式
	//3、权限0-7  0:没有任何权限 1:执行权限 2:写权限 3:写权限与执行权限 4:读权限 5:读权限与执行权限 6:读权限与写权限 7:读权限写权限执行权限
	//OpenFile不能创建新文件
	fp,err := os.OpenFile("D:/a.txt",os.O_RDWR,6)
	if err != nil{
		fmt.Println(err)
		return
	}
	//延时关闭
	//1、占用内存和缓冲区
	//2、文件打开上限pen
	defer fp.Close()

	//1、写入string文件
	n,_ := fp.WriteString("测试内容\r\n")
	fmt.Println(n)
	fp.WriteString("hehe")

	//2、写入byte切片
	var slice27 []byte
	slice27 = append(slice27,'a','b','c','d')
	fp.Write(slice27)
	//字符串和字符切片允许转换
	slice28 := []byte("hello world")
	fp.Write(slice28)
	//汉字转成3个字符
	slice29 := []byte("呵呵")
	fp.Write(slice29)

	//3、在文件制定位置写入
	//使用WriteAt制定位置插入数据会覆盖原来的内容
	//fp.WriteAt([]byte("呵呵呵水电费水电费递四方速递"),0)

	//移动游标
	//三个参数：1、偏移量 负数是向左偏移，正数是向右偏移
	//2、从哪里偏移 io.SeekCurrent 1  SeekEnd 2  SeekStart 0
	se,_ := fp.Seek(0,io.SeekEnd)
	fp.WriteAt([]byte("风湿热无热无热无\r\n"),se)
}