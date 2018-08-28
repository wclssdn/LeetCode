package main

import "fmt"

func main() {
	nums1 := []int{4, 5, 6, 0, 0, 0}
	nums2 := []int{1, 2, 3}
	merge(nums1, len(nums1)-len(nums2), nums2, len(nums2))

	fmt.Println(nums1)
}

//func merge(nums1 []int, m int, nums2 []int, n int) {
//	r := make([]int, m+n)
//	x := 0
//	y := 0
//	for i := 0; i < len(r); i++ {
//		if x < m && y < n {
//			if nums1[x] < nums2[y] {
//				r[i] = nums1[x]
//				x++
//			} else {
//				r[i] = nums2[y]
//				y++
//			}
//		} else if y < n {
//			r[i] = nums2[y]
//			y++
//		} else {
//			r[i] = nums1[x]
//			x++
//		}
//	}
//	for i := 0; i < len(r); i++ {
//		nums1[i] = r[i]
//	}
//}

func merge(nums1 []int, m int, nums2 []int, n int) {
	for i := len(nums1) - 1; i >= 0; i-- {
		if m > 0 && n > 0 {
			if nums1[m-1] > nums2[n-1] {
				m--
				nums1[i] = nums1[m]
			} else {
				n--
				nums1[i] = nums2[n]
			}
		} else if m > 0 {
			m--
			nums1[i] = nums1[m]
		} else {
			n--
			nums1[i] = nums2[n]
		}
	}
}
