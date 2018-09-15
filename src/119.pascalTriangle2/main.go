package main

import "fmt"

func main() {
	x := 5

	r := getRow(x)

	fmt.Println(r)
}
func getRow(rowIndex int) []int {
	result := make([]int, rowIndex+1)

	result[0] = 1

	for i := 0; i < rowIndex; i++ {
		result[i+1] = result[i] * (rowIndex - i) / (i + 1)
	}
	return result
}
