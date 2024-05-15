package easy

import "strings"

func LengthOfLongestSubstring(s string) int {
	past := strings.Split(s, "")
	if s == "" {
		return 0
	}
	maxSubstr := 1

	for len(past) > 1 {
		current := make([]string, 0)
		for i := 0; i < len(past)-1; i++ {
			if past[i+1] == "" {
				continue
			}
			if past[i] == "" {
				current = append(current, "")
				continue
			}
			a := []rune(past[i])
			b := []rune(past[i+1])
			if a[0] != b[len(b)-1] {
				current = append(current, past[i]+string(b[len(b)-1]))
				maxSubstr = len(past[i]) + 1
			} else {
				if len(current) == 1 {
					current = make([]string, 0)
				} else {
					current = append(current, "")
				}
			}
		}
		past = current
	}

	return maxSubstr
}
