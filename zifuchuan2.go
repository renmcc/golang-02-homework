package main

import "fmt"
import "unicode/utf8"

func main() {
	s := "Yes 啧啧啧!"
	//打印字符串内容
	fmt.Println(s)
	//打印字符串长度,计算的是byte长度
	fmt.Println(len(s))
	//打印16进制byte
	for _,b := range []byte(s) {
		fmt.Printf("%X ", b)
	}
	// ch is rune
	fmt.Println()
	for i, ch := range s {
		fmt.Printf("(%d %X)", i, ch)
	}
	//打印字符的长度
	fmt.Println()
	fmt.Println("Rune count:", utf8.RuneCountInString(s))
	//另一种方式打印字符串长度
	fmt.Println(len([]rune(s)))

	//遍历底层byte
	bytes := []byte(s)
	for len(bytes) > 0 {
		//ch 字节内容，size字节长度
		ch, size := utf8.DecodeRune(bytes)
		//进行切片
		bytes = bytes[size:]
		fmt.Printf("%c",ch)
	}

	//直接用runle类型去遍历字符
	fmt.Println()
	for _,ch := range []rune(s) {
		fmt.Printf("%c", ch)
	}

}
