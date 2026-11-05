import "slices"
func threeSum(nums []int) [][]int {
    res := make([][]int,0)
    slices.Sort(nums)
    for idx, num := range nums {
        if num > 0 {
            break
        }
        if idx > 0 && num == nums[idx-1] {
            continue
        }
        req := -num
        i := idx + 1
        j := len(nums) - 1
        for i < j {
            temp := nums[i] + nums[j]
            if (temp == req) {
                res = append(res,[]int{num,nums[i],nums[j]})
                i++
                j--
                for i < j && nums[i] == nums[i-1] { i++ }
                for i < j && nums[j] == nums[j+1] { j-- }
            } else if (temp > req) {
                j--
            } else {
                i++
            }
        }
    }
    return res
}
