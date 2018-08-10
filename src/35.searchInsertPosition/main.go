package main

import "fmt"

func main() {
	nums := []int{1, 3, 4, 6, 6}
	target := 2

	r := searchInsert(nums, target)

	fmt.Println("r", r)
}

func searchInsert(nums []int, target int) int {
	l := len(nums)
	half := l / 2
	if target < nums[half] {
		for i := 0; i < half; i++ {
			if target <= nums[i] {
				return i
			}
		}
		return half
	} else if target > nums[half] {
		for i := half; i < l; i++ {
			if target <= nums[i] {
				return i
			}
		}
		return l
	} else {
		for i := half; i > 0; i-- {
			if nums[i] < target {
				return i + 1
			}
		}
		return half
	}
}

//
//func searchInsert(nums []int, target int) int {
//	l := len(nums)
//	half := l / 2
//	if half == 1 {
//		return 1
//	}
//	if target < nums[half] {
//		return half + searchInsert(nums[0:half], target)
//	} else if target > nums[half] {
//		return half + searchInsert(nums[half:], target)
//	} else {
//		for i := half; i > 0; i-- {
//			if nums[i] < target {
//				return i + 1
//			}
//		}
//		return half
//	}
//}
