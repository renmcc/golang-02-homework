package main

import "fmt"

func main() {
	var x = map[string] string {
		"A": "Apple",
		"B": "Banana",
		"O": "Orange",
	}

	for key , val := range x {
		fmt.Println("Key:", key, "Value:", val)
	}
}
