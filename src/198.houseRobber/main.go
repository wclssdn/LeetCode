package main

import "fmt"

func main() {
	var nums = []int{1, 2, 3, 4, 5, 6}

	i := rob(nums)
	fmt.Println("rob:", i)
}

func rob(nums []int) int {
	prev1, prev2 := 0, 0
	for _, num := range nums {
		tmp := prev1
		if prev2+num > prev1 {
			prev1 = prev2 + num
		}
		prev2 = tmp
	}
	return prev1
}
