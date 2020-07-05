package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	rotate(nums, 3)
	fmt.Println("nums:", nums)

	nums = []int{1, 2, 3, 4, 5, 6}
	rotate(nums, 4)
	fmt.Println("nums:", nums)
}

func rotate(nums []int, k int) {
	rotate3(nums, k)
}

func rotate1(nums []int, k int) {
	for i := 0; i < k; i++ {
		tmp := nums[len(nums)-1]
		for j := len(nums) - 1; j > 0; j-- {
			nums[j] = nums[j-1]
		}
		nums[0] = tmp
	}
}

func rotate3(nums []int, k int) {
	if k == 0 {
		return
	}
	if k > len(nums) {
		k %= len(nums)
	}
	l := len(nums)
	xx := 0
	for i := 0; xx < l; i++ {
		x := nums[i]
		dest := i
		for {
			dest += k
			if dest >= l {
				dest -= l
			}
			x, nums[dest] = nums[dest], x
			xx++
			if dest == i {
				break
			}
		}
	}
}
