package main

import (
	"fmt"
	"strconv"
)

func main() {
	n := 5
	r := countAndSay(n)
	fmt.Println("N:", n, " say:", r)
}

func countAndSay(n int) string {
	if n == 1 {
		return "1"
	}
	str := countAndSay(n - 1)
	result := ""
	x := 1
	c := str[0]
	if len(str) == 1 {
		return "11"
	}
	for i := 1; i < len(str); i++ {
		if c == str[i] {
			x++
		} else {
			result += strconv.Itoa(x) + string(c)
			c = str[i]
			x = 1
		}
	}
	result += strconv.Itoa(x) + string(c)
	return result
}
