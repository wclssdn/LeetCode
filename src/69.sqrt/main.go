package main

import (
	"fmt"
)

func main() {
	x := 2147395600
	r := mySqrt(x)
	fmt.Printf("sqrt(%d) is %d", x, r)
}

func mySqrt(x int) int {
	if x <= 1 {
		return x
	}
	guess := float64(x) / 2
	tolerance := 0.1
	for {
		guess1 := (guess + float64(x)/guess) / 2
		if guess-guess1 <= tolerance {
			break
		}
		guess = guess1
	}
	if int(guess)*int(guess) > x {
		return int(guess - 1)
	}
	return int(guess)
}
