package medium

func lengthOfLongestSubstring(s string) int {
	past := []rune(s)
	charSet := make(map[rune]int)
	maxLength := 0
	currML := 0
	l := 0
	for i := 0; i < len(past); i++ {
		if _, ok := charSet[past[i]]; ok {
			tempL := charSet[past[i]] + 1
			if tempL > l {
				l = tempL
			}
		}
		currML = i - l + 1
		charSet[past[i]] = i
		maxLength = max(maxLength, currML)
	}
	return maxLength
}
