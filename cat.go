package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func printFile(name string) {
	buf, err := ioutil.ReadFile(name)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(buf))
}

func main() {
	if len(os.Args) < 2 {
		const Usage = "Usage: ./cat filename"
		fmt.Println(Usage)
		return
	}
	filename := os.Args[1]
	printFile(filename)
}