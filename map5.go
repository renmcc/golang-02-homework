package main

import "fmt"

func main() {
	var facebook = make(map[string]map[string]int)

	facebook["000001"] = map[string]int{"Jemy": 21}
	facebook["000002"] = map[string]int{"Tom": 25}

	for stu_no, stu_info := range facebook {
		fmt.Println("Student:", stu_no)

		for name,age := range stu_info {
			fmt.Println("Name:", name, "Age:", age)
		}
		fmt.Println()
	}
}
