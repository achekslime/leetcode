package solution

import (
	"fmt"
	"strings"
)

func LongestPalindrome(s string) []int {
	ss := strings.Split(s, "")
	d := make([]int, len(ss))
	l, r := 0, -1

	for i := range d {
		if i > r {
			v := Primitive(ss, d, i)
			if v > 1 {
				delta := (v - 1) / 2
				l = i - delta
				r = i + delta
			}
		} else {
			if i == 7 {
				fmt.Println("lol")
			}
			deltaFromEdge := r - i
			reflectedValue := d[l+deltaFromEdge]

			rNew := i + (reflectedValue-1)/2
			deltaOverEnd := rNew - (len(ss) - 1)
			if deltaOverEnd > 0 {
				reflectedValue -= deltaOverEnd * 2
			}
			d[i] = reflectedValue
			v := Primitive(ss, d, i)
			if v > 1 {
				delta := (v - 1) / 2
				if i+delta > r {
					l = i - delta
					r = i + delta
				}
			}
		}
	}

	return d
}

func Primitive(s []string, d []int, i int) int {
	delta := 1
	if d[i] != 0 {
		delta = (d[i] - 1) / 2
	} else {
		d[i] = 1
	}

	for i-delta >= 0 && i+delta < len(s) {
		if s[i-delta] == s[i+delta] {
			d[i] = delta*2 + 1
			delta += 1
		} else {
			break
		}
	}
	return d[i]
}
