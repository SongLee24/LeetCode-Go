package main

// 方法一：a 和 b 两个数组的头部分别维护两个指针，若其中一个比另一个小，则向前移动，若遇到相等时保存，遍历直到其中一个数组的尾部，时间复杂度O(m+n)；
func intersection(arr1, arr2 []int) []int {
	resArr := make([]int, 0)
	if len(arr1) == 0 || len(arr2) == 0 {
		return resArr
	}
	arr1Index := 0
	arr2Index := 0

	for arr1Index < len(arr1) && arr2Index < len(arr2) {
		if arr1[arr1Index] == arr2[arr2Index] {
			resArr = append(resArr, arr1[arr1Index])
			arr1Index += 1
			arr2Index += 1
		} else if arr1[arr1Index] < arr2[arr2Index] {
			arr1Index += 1
		} else {
			arr2Index += 1
		}
	}
	return resArr
}

// 方法二：将 a 中的元素 hash 存储（用map或者dict），遍历 b 中的每一个值看是否在这个hash 中，若存在就保存，时间复杂度是 O(m)，空间复杂度是O(n)；
