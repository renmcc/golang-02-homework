package main

import (
	"fmt"
	"time"
)

var abcd int = 1

func main() {
	for {
		fmt.Printf("%p\n", &abcd)
		fmt.Println(abcd)
		time.Sleep(3*time.Second)
	}
}