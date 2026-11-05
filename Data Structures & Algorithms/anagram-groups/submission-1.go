import "slices"
func groupAnagrams(strs []string) [][]string {
    res := make([][]string,0)
    seen := make(map[string][]string)
    for _, str := range strs {
        runes := []rune(str)
        slices.Sort(runes)
        sortedStr := string(runes)
        seen[sortedStr] = append(seen[sortedStr],str)
    }
    for _, v := range seen {
        res = append(res,v)
    }
    return res
}
