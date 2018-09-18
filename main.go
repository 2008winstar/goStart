package main

import "fmt"

func main() {
	x := 1
	for x <= 10 {
		if x%2 == 0 {
			fmt.Println(x, "even")
		} else {
			fmt.Println(x, "odd")
		}
		x += 1
	}
}
