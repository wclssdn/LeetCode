package main

import "fmt"

// https://leetcode.com/problems/majority-element/
func main() {
	ints := []int{1, 2, 3, 4, 3, 2, 1, 3, 3, 2, 3, 2, 3, 3}
	element := majorityElement(ints)
	fmt.Println("majority element:", element)
}

func majorityElement(nums []int) int {
	maxTimes := 0
	maxTimesElement := 0
	m := make(map[int]int, len(nums)/2+1)
	for _, v := range nums {
		m[v]++
		if m[v] > maxTimes {
			maxTimes = m[v]
			maxTimesElement = v
			if maxTimes > len(nums)/2 {
				return maxTimesElement
			}
		}
	}
	return maxTimesElement
}
