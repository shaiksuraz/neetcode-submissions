func validPalindrome(s string) bool {
	isPalindrome := func(l, r int) bool {
		for l < r {
			if s[l] != s[r] {
				return false
			}
			l++
			r--
		}
		return true
	}
	i := 0
	j := len(s) - 1
	for i < j {
		if s[i] != s[j] {
			return isPalindrome(i+1, j) || isPalindrome(i, j-1)
		}
		i++
		j--
	}
	return true
}