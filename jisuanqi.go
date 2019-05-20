package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	stra := os.Args[1]
	op := os.Args[2]
	strb := os.Args[3]

	var a, b int
	var err error
	if a, err = strconv.Atoi(stra); err != nil {
		log.Fatal(err)
	}
	if b, err = strconv.Atoi(strb); err != nil {
		log.Fatal(err)
	}

	var ret int
	switch  op {
	case "+":
		ret = a + b
	case "-":
		ret = a - b
	case "*":
		ret = a * b
	case "/":
		ret = a / b
	default:
		fmt.Println("Invalid op" + op)
	}
	fmt.Println(ret)
}
