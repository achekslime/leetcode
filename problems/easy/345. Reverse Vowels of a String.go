package easy

func reverseVowels(s string) string {
	ss := []rune(s)

	vowels := make([]rune, 0)
	positions := make([]int, 0)

	for i, v := range ss {
		if v == 'a' || v == 'e' || v == 'i' || v == 'o' || v == 'u' {
			vowels = append(vowels, v)
			positions = append(positions, i)
			continue
		}
		if v == 'A' || v == 'E' || v == 'I' || v == 'O' || v == 'U' {
			vowels = append(vowels, v)
			positions = append(positions, i)
		}
	}

	lastPosition := len(positions) - 1
	for i := lastPosition; i >= 0; i-- {
		ss[positions[i]] = vowels[lastPosition-i]
	}

	return string(ss)
}
