package main

import "fmt"

func main() {
	nums := []int{1, 7, 3, 6, 5, 6}
	index := pivotIndex(nums)
	fmt.Println(index)
}

func pivotIndex(nums []int) int {
	leftSum := 0
	rightSum := 0
	for _, num := range nums {
		rightSum += num
	}

	for i, num := range nums {
		rightSum -= num
		if leftSum == rightSum {
			return i
		}
		leftSum += num
	}
	return -1
}
