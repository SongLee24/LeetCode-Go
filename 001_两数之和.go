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

