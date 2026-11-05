// import "math"
func longestConsecutive(nums []int) int {
    set := make(map[int]struct{})
    for _, num := range nums {
        set[num] = struct{}{}
    }
    var res int
    for _, num := range nums {
        //check left exists
        if _, ok := set[num-1]; !ok {
            //sequence starts here
            //go until we find right
            var temp int = 1
            for {
                if _, ok := set[num+1]; !ok {
                    break
                } else {
                    temp++
                    num++
                }
            }
            res = max(res,temp)
        } else {
            continue
        }
    }
    return res
}
