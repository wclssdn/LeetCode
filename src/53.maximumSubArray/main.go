package main

import "fmt"

func main() {
	nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	//nums = []int{-1, 0, -2}
	r := maxSubArray(nums)
	fmt.Println("r", r)
}

//func maxSubArray(nums []int) int {
//	max := nums[0]
//	for i := 0; i < len(nums); i++ {
//		tmp := nums[i]
//		if max < tmp {
//			max = tmp
//		}
//		for j := i + 1; j < len(nums); j++ {
//			tmp += nums[j]
//			if max < tmp {
//				max = tmp
//			}
//		}
//
//	}
//	return max
//}

func maxSubArray(nums []int) int {
	max := nums[0]
	if len(nums) == 1 {
		return max
	}
	tmp := max
	for i := 1; i < len(nums); i++ {
		// hinge
		if tmp+nums[i] > nums[i] {
			tmp += nums[i]
		} else {
			tmp = nums[i]
		}
		if max < tmp {
			max = tmp
		}
	}
	return max
}
