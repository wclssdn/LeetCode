package main

import "fmt"

func main() {
	x := 21120
	r := isPalindrome(x)
	fmt.Println("r:", r)
}

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	if x < 10 {
		return true
	}
	if x%10 == 0 {
		return false
	}

	pop := 0
	for {
		if x < 10 {
			break
		}
		pop = pop*10 + x%10

		if pop >= x/10 {
			x /= 10
			if pop > x {
				pop /= 10
			}
			break
		}
		x /= 10
	}

	return pop == x
}
