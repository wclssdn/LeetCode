package main

import "fmt"

func main() {
	var x uint32 = 937919
	bits := hammingWeight(x)
	fmt.Printf("num: %032b\n", x)
	fmt.Println("bits:", bits)
}

func hammingWeight(num uint32) int {
	result := 0
	for i := 0; i < 32; i++ {
		if num>>i&1 == 1 {
			result++
		}
	}
	return result
}

// num-1 change the latest "1" to "0"
// after num = num & (num-1) the latest "1" and all bits after it was change to "0"
func resolve(num uint32) int {
	count := 0
	for num != 0 {
		count++
		num = num & (num - 1)
	}
	return count
}
