package main

import "fmt"

type rect struct {
	width, height int
}

func (r *rect) area() int {
	return r.width * r.height
}

func (r rect) perim() int {
	return 2*r.width + 2*r.height
}


func main() {
	b := rect{width:10, height:5}
	fmt.Println(b.area())
	fmt.Println(b.perim())

}