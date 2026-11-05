func isAlphaNumeric(r rune) bool {
	return unicode.IsLetter(r) || unicode.IsDigit(r)
}

func isPalindrome(s string) bool {
	runes := []rune(s)
	i := 0
	j := len(runes) - 1
	for i < j {
		if !isAlphaNumeric(runes[i]) {
			i++
			continue
		}
		if !isAlphaNumeric(runes[j]) {
			j--
			continue
		}
		if unicode.ToLower(runes[i]) != unicode.ToLower(runes[j]) {
			return false
		}
		i++
		j--
	}
	return true
}
