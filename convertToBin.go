package main

import (
	"fmt"
	"strconv"
)

func convertToBin(n int) string {
	result := ""
	for ; n > 0; n /=2 {
		lsb := n % 2
		fmt.Println(lsb)
		result = strconv.Itoa(lsb) + result
	}
	return result
}

func main() {
	fmt.Println(convertToBin(128))
	fmt.Println(convertToBin(192))
}
