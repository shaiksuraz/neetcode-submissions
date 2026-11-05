import "slices"
func groupAnagrams(strs []string) [][]string {
    res := make([][]string,0)
    seen := make(map[string][]string)
    for _, str := range strs {
        runes := []rune(str)
        slices.Sort(runes)
        sortedStr := string(runes)
        if _, ok := seen[sortedStr]; !ok {
            seen[sortedStr] = make([]string,0)
        }
        seen[sortedStr] = append(seen[sortedStr],str)
    }
    fmt.Println(seen)
    for _, v := range seen {
        res = append(res,v)
    }
    return res
}
