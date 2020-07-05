package main

import "fmt"

func main() {
	var x uint32 = 123456789
	fmt.Printf("x's bits:\t %032b\n", x)
	reversed := reverseBits(x)
	fmt.Printf("reversed bits:\t %032b\n", reversed)
}

func reverseBits(num uint32) uint32 {
	var reversed uint32 = 0
	var i uint32 = 0
	for ; i < 32; i++ {
		//rightBit := ((1 << i) & num) >> i
		//leftBit := ((1 << (31 - i)) & num) >> (31 - i)
		//t := (leftBit ^ rightBit) << (31 - i)
		//reversed |= t
		reversed |= ((((1 << i) & num) >> i) ^ (((1 << (31 - i)) & num) >> (31 - i))) << (31 - i)
	}
	return reversed ^ num
}
