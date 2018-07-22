package main

import "fmt"

func main() {
	a := []int{-3, 1, 3, 90}
	target := 0
	fmt.Println("indeies:", twoSum(a, target))
}

// 48 ms  top 31%
//func twoSum(nums []int, target int) []int {
//	i, j := 0, 1
//find:
//	for i = 0; i < len(nums); i++ {
//		for j = i + 1; j < len(nums); j++ {
//			if nums[i]+nums[j] == target {
//				break find
//			}
//		}
//	}
//	return []int{i, j}
//}

// 4ms/8ms top 1
// 利用hash表存储待查找的具体值
// 8ms的解先预生成了所有元素的hash表。。。这个就是正向思维。。。
func twoSum(nums []int, target int) []int {
	h := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		j := target - nums[i]
		if _, ok := h[j]; ok {
			return []int{h[j], i}
		}
		h[nums[i]] = i
	}
	return nil
}
