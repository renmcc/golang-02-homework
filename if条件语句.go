package main

import "fmt"

func main() {
	var score int
	fmt.Scan(&score)
	if score > 700 {
		fmt.Println("大于700")
	}else if score > 600 {
		fmt.Println("大于600")
	}else{
		fmt.Println("小于600")
	}
}