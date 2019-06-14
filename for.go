package main

import "fmt"

func main() {
	s := "世界hello"
	s1 := []rune(s)
	for i := 0; i < len(s1); i++ {
		fmt.Println(string(s1[i]))
	}

	var arr = [...]int{1,2,3}
	fmt.Printf("%p\n", &arr)
	fmt.Printf("%p\n", &arr[0])

	const a = 10
	fmt.Printf("%T", a)
}
