func hasDuplicate(nums []int) bool {
    freq:= make(map[int]bool,len(nums))
    for _, num := range nums {
        if _, ok := freq[num]; !ok {
            freq[num] = true
        } else {
            return true
        }
    }
    return false
}
