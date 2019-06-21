package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "hello world"

	//判断字符串是否包含，返回布尔值
	value := strings.Contains(str, "wo")
	fmt.Println(value)

	//切片转字符串
	s := []string{"123","456","789"}
	str1 := strings.Join(s, "-")
	fmt.Println(str1)

	//找匹配字符的下标,返回int 找不到返回-1
	str2 := "日照香炉生紫烟"
	value2 := strings.Index(str2,"生紫烟")
	fmt.Println(value2)

	//重复字符多少次 count 要大于0 0什么不显示
	str3 := "你愁啥"
	value3 := strings.Repeat(str3, 10)
	fmt.Println(value3)

	//字符串替换 n -1表示全部替换
	str4 := "hello world"
	value4 := strings.Replace(str4, "l","2",2)
	fmt.Println(value4)

	//字符串转切片
	str5 := "www.baidu.com"
	value5 := strings.Split(str5, ".")
	fmt.Println(value5)

	//去掉字符串首尾指定字符
	str6 := " hello world "
	value6 := strings.Trim(str6, " ")
	fmt.Println(value6)

	//去掉字符串中的空格，并按照空格分隔，返回切片
	str7 := "  are u  ok ?   "
	value7 := strings.Fields(str7)
	fmt.Println(value7)

}