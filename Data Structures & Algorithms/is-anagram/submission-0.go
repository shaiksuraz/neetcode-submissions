func isAnagram(s string, t string) bool {
    if len(s) != len(t) {
        return false
    }
    sMap := make(map[rune]int)
    tMap := make(map[rune]int)
    for _, r := range s {
        if _, ok := sMap[r]; !ok {
            sMap[r] = 1
        } else {
            sMap[r]++
        }
    }
    for _, r := range t {
        if _, ok := tMap[r]; !ok {
            tMap[r] = 1
        } else {
            tMap[r]++
        }
    }
    if len(sMap) != len(tMap) {
        return false
    }
    for k1, v1 := range sMap {
        if v2, ok := tMap[k1]; !ok || v1!=v2 {
            return false
        }
    }
    return true
}
