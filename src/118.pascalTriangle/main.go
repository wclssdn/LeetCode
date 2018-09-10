package main

import "fmt"

func main() {
	x := 0

	r := generate(x)

	fmt.Println(r)
}

func generate(numRows int) [][]int {
	var result [][]int
	if numRows == 0 {
		return result
	}
	for i := 0; i < numRows; i++ {
		result = append(result, make([]int, i+1))
	}

	result[0][0] = 1
	if numRows == 1 {
		return result
	}

	result[1][0] = 1
	result[1][1] = 1
	for i := 2; i < numRows; i++ {
		tmp := make([]int, i+1)
		tmp[0] = 1
		tmp[i] = 1
		for j := 1; j < i; j++ {
			tmp[j] = result[i-1][j-1] + result[i-1][j]
		}
		for k, v := range tmp {
			result[i][k] = v
		}
	}
	return result
}
