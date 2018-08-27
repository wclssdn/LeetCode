package main

import "fmt"

func main() {
	n := 7
	r := climbStairs(n)
	fmt.Println("r:", r)
}

func climbStairs(n int) int {
	if n <= 2 {
		return n
	}

	r := 0
	f1 := 2
	f2 := 1
	for i := 3; i <= n; i++ {
		r = f1 + f2
		f2 = f1
		f1 = r
	}
	return r
}
