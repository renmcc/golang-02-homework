package main

import "fmt"

func main() {
	var arr1 = [...]int{1,2,23,4,5}
	var s1 = arr1[2:3]
	var s2 = arr1[:3]
	var s3 = arr1[2:]
	var s4 = arr1[:]

	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(s3)
	fmt.Println(s4)
}
