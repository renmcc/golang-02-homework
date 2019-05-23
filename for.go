package main

import "fmt"

func main() {
	s := "世界hello"
	s1 := []rune(s)
	for i := 0; i < len(s1); i++ {
		fmt.Println(string(s1[i]))
	}

}
