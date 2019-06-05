package main

import "fmt"

func zhizhen(i int) {
	i = 10
}

func zhizhen2(i *int) {
	*i = 10
}

func main() {
	i := 1
	fmt.Println(i)
	zhizhen(i)
	fmt.Println(i)
	zhizhen2(&i)
	fmt.Println(i)
}