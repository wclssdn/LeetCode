package main

import "fmt"

func main() {
	nums := []int{0, 1, 2, 2, 3, 0, 4, 2}

	r := removeElement(nums, 2)
	fmt.Println(nums, r)
}

func removeElement(nums []int, val int) int {
	pos := 0
	length := len(nums)
	for i := 0; i < length; i++ {
		if val == nums[pos] {
			nums = append(nums[:pos], nums[pos+1:]...)
		} else {
			pos++
		}
	}
	return len(nums)
}
