import "slices"
func topKFrequent(nums []int, k int) []int {
    freq := make(map[int]int)
    for _, num := range nums {
        freq[num]++
    }
    fmt.Println(freq)
    arr := make([][2]int, 0 ,len(freq))
    for num, cnt := range freq {
        arr = append(arr,[2]int{cnt, num})
    }
    slices.SortFunc(arr, func(i,j [2]int) int {
        return j[0] - i[0]
    })
    res := make([]int, k)
    for i:=0 ; i<k; i++ {
        res[i] = arr[i][1]
    }
    return res
}
