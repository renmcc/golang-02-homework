package main

import "fmt"
import "os"
import "strconv"

func grade(score int) string {
	g := ""
	switch {
		case score < 0 || score > 100:
			panic(fmt.Sprintf("Wrong score: %d", score))
		case score < 60:
			g = "F"
		case score < 80:
			g = "C"
		case score <= 100:
			g = "A"
		default:
			g = "Z"
	}
	return g
}

func main() {
	args,_ := strconv.Atoi(os.Args[1])
	score := grade(args)
	fmt.Println(score)
}
