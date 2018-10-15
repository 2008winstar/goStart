package main

import "fmt"

func getGDC(x, y int) int {
	if y == 0 {
		return x
	}
	r := x % y
	return getGDC(y, r)
}

func main() {
	fmt.Println(getGDC(126, 90))
}
