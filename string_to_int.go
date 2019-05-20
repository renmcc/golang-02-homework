package main

import (
	"fmt"
	"log"
	"strconv"
)

func main() {
	s := "12"
	x, err := strconv.Atoi(s)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("x=", x)
}
