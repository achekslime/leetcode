package medium

import (
	"strings"
)

func longestPalindrome(s string) string {
	ss := strings.Split(s, "")
	maxOdd := OddPalindromes(ss)
	maxEven := EvenPalindromes(ss)
	if len(maxOdd) > len(maxEven) {
		return maxOdd
	} else {
		return maxEven
	}
}

func EvenPalindromes(s []string) string {
	d := make([]int, len(s))
	l, r := 0, -1
	maxValue := ""

	for i := 0; i < len(s)-1; i++ {
		var v int
		if i+1 > r {
			v = EvenPrimitive(s, d, i)
		} else {
			deltaFromRightEdge := r - i - 2
			if deltaFromRightEdge >= 0 {
				reflectedValue := d[l+deltaFromRightEdge]
				rImaginary := i + (reflectedValue-2)/2 + 1
				overflow := rImaginary - (r - 1)
				if overflow > 0 {
					reflectedValue -= overflow * 2
				}
				d[i] = reflectedValue
			}
			v = EvenPrimitive(s, d, i)
		}

		delta := (v - 2) / 2
		if i+delta+1 > r {
			l = i - delta
			r = i + delta + 1
		}
		if v > len(maxValue) {
			maxValue = strings.Join(s[i-delta:i-delta+v], "")
		}
	}

	return maxValue
}

func EvenPrimitive(s []string, d []int, i int) int {
	delta := 0
	if d[i] > 1 {
		delta = (d[i] - 1) / 2
	} else {
		d[i] = 0
	}

	for i-delta >= 0 && i+1+delta < len(s) {
		if s[i-delta] == s[i+1+delta] {
			d[i] = delta*2 + 2
			delta += 1
		} else {
			break
		}
	}

	return d[i]
}

func OddPalindromes(s []string) string {
	d := make([]int, len(s))
	l, r := 0, -1
	maxValue := ""

	for i := range d {
		var v int
		if i > r {
			v = Primitive(s, d, i)
		} else {
			deltaFromRightEdge := r - i
			reflectedValue := d[l+deltaFromRightEdge]
			rImaginary := i + reflectedValue/2
			overflow := rImaginary - r
			if overflow > 0 {
				reflectedValue -= overflow * 2
			}
			d[i] = reflectedValue
			v = Primitive(s, d, i)
		}

		delta := v / 2
		if i+delta > r {
			l = i - delta
			r = i + delta
		}
		if v > len(maxValue) {
			maxValue = strings.Join(s[i-delta:i-delta+v], "")
		}
	}

	return maxValue
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
