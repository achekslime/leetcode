package easy

func mergeAlternately(word1 string, word2 string) string {
	a := []rune(word1)
	b := []rune(word2)
	n, m := len(a), len(b)
	i, j, k := 0, 0, 0
	result := make([]rune, n+m)
	for i < n || j < m {
		if i < n {
			result[k] = a[i]
			i++
			k++
		}
		if j < m {
			result[k] = b[j]
			j++
			k++
		}
	}
	return string(result)
}
