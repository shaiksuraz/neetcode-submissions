func isAnagram(s string, t string) bool {
    if len(s) != len(t) {
        return false
    }
    sMap := make(map[rune]int)
    for _, r := range s {
            sMap[r]++
    }
    for _, r := range t {
            sMap[r]--
            if(sMap[r] <0) {
                return false
            }
    }
    return true
}
