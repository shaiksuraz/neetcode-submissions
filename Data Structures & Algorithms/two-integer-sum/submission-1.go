func twoSum(nums []int, target int) []int {
    seen := make(map[int]int)
    for id1, num := range nums {
        if id2, ok := seen[target-num]; ok {
            return []int{id2, id1}
        }
           seen[num] = id1
    }
    return nil
}
