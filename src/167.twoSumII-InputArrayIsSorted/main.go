package main

import (
	"fmt"
)

func main() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 20, 30, 40, 50, 60}
	result := twoSum(numbers, 15)
	fmt.Println("result:", result)
}

func twoSum(numbers []int, target int) []int {
	i := 0
	j := len(numbers) - 1
	for {
		sum := numbers[i] + numbers[j]
		if sum > target {
			j--
		} else if sum < target {
			i++
		} else {
			return []int{i + 1, j + 1}
		}
	}
}
