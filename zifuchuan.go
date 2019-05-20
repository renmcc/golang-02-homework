package main

import "fmt"

//翻转字符串
func reverse(s string) string {
	s1 := []rune(s)
	s2 := make([]rune, len(s1))
	for i := 0; i < len(s1); i++ {
		s2[i] = s1[len(s1)-i-1]
	}
	return string(s2)
}

func main() {
	//字符串不可变，需要修改成byte再修改
	s := "hello"
	s1 := []byte(s)
	s1[0] = 'a'
	s = string(s1)
	fmt.Println(s)

	//另一种方式，切片
	s = "h" + s[1:]
	fmt.Println(s)

	//如果是汉字需要转换成rune类型
	s = "你好，世界"
	s2 := []rune(s)
	s2[0] = 'a'
	s = string(s2)
	fmt.Println(s)
	fmt.Println(reverse("abcd世界"))
}
