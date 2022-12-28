package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4}
	result := runningSum2(nums)
	fmt.Println(nums, result)
}

func runningSum(nums []int) []int {
	lastSum := 0
	result := make([]int, len(nums))
	for i, num := range nums {
		result[i] = lastSum + num
		lastSum += num
	}
	return result
}

func runningSum2(nums []int) []int {
	lastSum := 0
	for i, num := range nums {
		nums[i] = lastSum + num
		lastSum += num
	}
	return nums
}
