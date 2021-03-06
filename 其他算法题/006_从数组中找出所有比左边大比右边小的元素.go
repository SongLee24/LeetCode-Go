package main

import "fmt"

func GetPivotElements(arr []int) []int {
	length := len(arr)
	if length < 3 {
		return []int{}
	}

	minIndex := make([]int, length)
	min := arr[length-1]
	for i := length - 1; i >= 0; i-- {
		if arr[i] < min {
			min = arr[i]
		}
		minIndex[i] = min
	}

	var ret []int
	for i := 1; i < length-1; i++ {
		if arr[i] > arr[i-1] && arr[i] < minIndex[i+1] {
			ret = append(ret, arr[i])
		}
	}
	return ret
}

func main() {
	testArray := []int{1, 8, 6, 9, 10, 15, 12, 20} // minIndex: [1 6 6 9 10 12 12 20]
	fmt.Println(GetPivotElements(testArray))       // [9, 10]
}
