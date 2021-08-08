// 双循环
func twoSum(nums []int, target int) []int {
    var index1 int
    for i, v := range nums {
        index1 = i
        val := target - v
        for k := i+1; k < len(nums); k++ {
            if nums[k] == val {
                return []int{index1, k}
            }
        }
    }
    return []int{}
}

// 使用map缓存
func twoSum(nums []int, target int) []int {
    cache := make(map[int]int)
    for i, v := range nums {
        another := target - v
        if ii, ok := cache[another]; ok {
            return []int{i, ii}
        } else {
            cache[v] = i
        }
    }
    return []int{}
}

