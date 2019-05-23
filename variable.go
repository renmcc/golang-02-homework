package main

import "fmt"

func main01() {
	a := 5
	b := 3.141592
	fmt.Print("a")
	fmt.Print("b")
	fmt.Printf("==%05d==", a)
	fmt.Printf("%.4f", b)
}

func main() {
	var a int
	var b,c string
	fmt.Scan("%3d", &a)
	fmt.Scan("%s", &b)
	fmt.Scan("%s", &c)
	fmt.Println(a,b,c)

}
