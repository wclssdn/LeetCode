package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 9}
	arr = []int{9}
	r := plusOne(arr)
	fmt.Println(r)
}

func plusOne(digits []int) []int {
	carry := false
	for i := len(digits) - 1; i >= 0; i-- {
		if digits[i] == 9 {
			digits[i] = 0
			carry = true
		} else {
			digits[i]++
			carry = false
			break
		}
	}
	if carry == true {
		digits = append([]int{1}, digits...)
	}
	return digits
}
