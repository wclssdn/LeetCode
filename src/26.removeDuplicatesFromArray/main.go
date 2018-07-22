package main

import (
	"fmt"
)

/**
给定一个排序数组，你需要在原地删除重复出现的元素，使得每个元素只出现一次，返回移除后数组的新长度。

不要使用额外的数组空间，你必须在原地修改输入数组并在使用 O(1) 额外空间的条件下完成。

示例 1:

给定数组 nums = [1,1,2],

函数应该返回新的长度 2, 并且原数组 nums 的前两个元素被修改为 1, 2。

你不需要考虑数组中超出新长度后面的元素。
示例 2:

给定 nums = [0,0,1,1,1,2,2,3,3,4],

函数应该返回新的长度 5, 并且原数组 nums 的前五个元素被修改为 0, 1, 2, 3, 4。

你不需要考虑数组中超出新长度后面的元素。
说明:

为什么返回数值是整数，但输出的答案是数组呢?

请注意，输入数组是以“引用”方式传递的，这意味着在函数里修改输入数组对于调用者是可见的。

你可以想象内部操作如下:

// nums 是以“引用”方式传递的。也就是说，不对实参作任何拷贝
int len = removeDuplicates(nums);

// 在函数里修改输入数组对于调用者是可见的。
// 根据你的函数返回的长度, 它会打印出数组中该长度范围内的所有元素。
for (int i = 0; i < len; i++) {
    print(nums[i]);
}
*/
func main() {
	var nums = []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	l := removeDuplicates(nums)
	fmt.Println("len:", l, "nums:", nums)
}

// 方案1：判断重复，后边元素前移  161 / 161 个通过测试用例
//状态：通过
//执行用时：276 ms
//func removeDuplicates(nums []int) int {
//	numsLen := len(nums)
//	if numsLen == 0{
//		return 0
//	}
//	finalLen := numsLen;
//	current := nums[0]
//	for i:=1; i < numsLen; i++{
//		tmp := nums[i]
//		if nums[i] == current{
//			if (i == finalLen){
//				break;
//			}
//			finalLen--
//			for j := i; j < numsLen -1 ; j++ {
//				nums[j] = nums[j+1]
//			}
//			i--
//			continue
//		}
//
//		current = tmp
//	}
//	return finalLen
//}

// 方案2：判断是否重复，不重复则移动到当前排重顺序位置
func removeDuplicates(nums []int) int {
	l := len(nums)
	if l == 0 {
		return 0
	}

	p := 0
	for i := 1; i < l; i++ {
		if nums[p] != nums[i] {
			p++
			nums[p] = nums[i]
		}
	}
	return p + 1
}
