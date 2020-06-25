package main

import (
	"fmt"
)

func main() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 20, 30, 40, 50, 60}
	find, err := findOffset(numbers, 10)
	fmt.Println("find ", find, err)
	result := twoSum(numbers, 15)
	fmt.Println("result:", result)
}

func twoSum(numbers []int, target int) []int {
	for i, v := range numbers {
		offset, err := findOffset(numbers, target-v)
		if err == 0 {
			return []int{i + 1, offset + 1}
		}
	}
	return []int{}
}

func findOffset(numbers []int, target int) (int, int) {
	if len(numbers) == 0 {
		return 0, 1
	}
	if len(numbers) == 1 && numbers[0] != target {
		return 0, 1
	}
	half := len(numbers) / 2
	if numbers[half] == target {
		return half, 0
	} else if numbers[half] < target {
		offset, err := findOffset(numbers[half:], target)
		if err != 0 {
			return 0, err
		}
		return half + offset, 0
	} else {
		return findOffset(numbers[:half], target)
	}
}
