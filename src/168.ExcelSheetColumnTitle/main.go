package main

import "fmt"

func main() {
	for _, i := range []int{1, 2, 25, 26, 27, 28, 52, 53, 701, 26 * 26, 26 * 26 * 26} {
		fmt.Println("i:", i, "title:", convertToTitle(i))
	}
}

func convertToTitle(n int) string {
	result := make([]byte, 0, 10)
	for n != 0 {
		n--
		b := n % 26
		result = append(result, 'A'+byte(b))
		n = n / 26
	}
	// reserve
	r := make([]byte, 0, len(result))
	for i := len(result) - 1; i >= 0; i-- {
		r = append(r, result[i])
	}
	return string(r)
}
