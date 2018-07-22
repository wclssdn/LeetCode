package main

import (
	"fmt"
	"math"
)

func main() {
	x := 1563847412
	r := reverse(x)
	fmt.Println("reverse:", r)
}

func reverse(x int) int {
	navigate := false
	if x < 0 {
		navigate = true
		x = int(math.Abs(float64(x)))
	}
	r := 0
	for {
		pop := x % 10
		r = r*10 + pop

		if x < 10 {
			break
		}
		x /= 10
	}
	if float64(r) >= math.Pow(2, 31)-1 {
		return 0
	}
	if navigate {
		return -r
	}
	return r
}
