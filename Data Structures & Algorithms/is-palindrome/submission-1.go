func isPalindrome(s string) bool {
    var runes []rune
    for _ , r := range s {
        if unicode.IsDigit(r) || unicode.IsLetter(r) {
            runes = append(runes, unicode.ToUpper(r))
        }
    }
    // fmt.Println(runes)
    i := 0
    j := len(runes) - 1
    for i<j {
        // fmt.Println(runes[i], runes[j])
        if runes[i] != runes[j] {
            return false
        }
        i++
        j--
    }
    return true
}
