package main

import "fmt"

func reverses(s string) string {
	s1 := []rune(s)
	s2 := make([]rune, len(s1))
	for i:= 0; i< len(s1); i++ {
		s2[i] = s1[len(s1)-i-1]
	}
	return string(s2)
}

func main() {
	s := "你好，世界"
	s1 := reverses(s)
	fmt.Println(s1)
}
