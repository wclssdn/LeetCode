package main

import (
	"fmt"
	"math"
)

func main() {
	titles := []string{"A", "B", "Z", "AA", "AZ", "BA"}
	for _, title := range titles {
		fmt.Println("title", title, "number", titleToNumber(title))
	}
}

func titleToNumber(s string) int {
	bytes := []byte(s)
	result := 0
	for i := len(bytes) - 1; i >= 0; i-- {
		pow := float64(len(bytes) - i - 1)
		result += int(bytes[i]-'A'+1) * int(math.Pow(26, pow))
	}
	return result
}
