package main

import "fmt"

func main() {
	haystack := "hello"
	needle := "ll"

	r := strStr(haystack, needle)
	fmt.Printf("find [%s] in [%s], result: %d\n", needle, haystack, r)
}

//
//func strStr(haystack string, needle string) int {
//	if len(needle) == 0 {
//		return 0
//	}
//
//	result := -1
//	l := len(haystack)
//	l2 := len(needle)
//	for i := 0; i < l; i++ {
//		for j := 0; j < l2; j++ {
//			if i+j < l && needle[j] == haystack[i+j] {
//				if j == 0 {
//					result = i
//				}
//			} else {
//				result = -1
//				break
//			}
//		}
//		if result != -1 {
//			break
//		}
//	}
//	return result
//}

func strStr(haystack string, needle string) int {
	if len(needle) == 0 {
		return 0
	}

	result := -1
	l := len(haystack)
	l2 := len(needle)
	for i := 0; i < l; i++ {
		if l < i+l2 {
			return -1
		}
		if haystack[i:i+l2] == needle {
			return i
		}
	}
	return result
}
