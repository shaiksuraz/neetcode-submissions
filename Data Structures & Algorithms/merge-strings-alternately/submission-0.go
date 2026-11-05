func mergeAlternately(word1 string, word2 string) string {
	l1 := len(word1)
	l2 := len(word2)
	str := make([]byte, l1+l2)
	i := 0
	j := 0
	k := 0
	for i < l1 && j < l2 {
		str[k] = word1[i]
		k++
		i++
		str[k] = word2[j]
		k++
		j++
	}
	for i < l1 {
		str[k] = word1[i]
		k++
		i++
	}
	for j < l2 {
		str[k] = word2[j]
		k++
		j++
	}

	return string(str)
}
