package medium

import (
	"strings"
)

func longestPalindrome(s string) string {
	ss := ">#" + strings.Join(strings.Split(s, ""), "#") + "#<"
	n := len(ss)
	d := make([]int, n)
	c, r := 0, 0

	for i := 0; i < n-1; i++ {
		if i <= r {
			d[i] = min(d[2*c-i], r-i)
		}
		for j := i + 1; j < n-1; j++ {
			if 2*i-j > 0 && ss[2*i-j] == ss[j] {
				d[i]++
			} else {
				break
			}
		}
		if i+d[i] > r {
			c, r = i, d[i]
		}
	}

	maxLen, maxC := 0, 0
	for i, v := range d {
		if d[i] > maxLen {
			maxLen = v
			maxC = i
		}
	}
	return s[(maxC-maxLen)/2 : (maxC+maxLen)/2]
}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}
