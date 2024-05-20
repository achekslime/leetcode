package easy

func isHappy(n int) bool {
	mp := make(map[int]bool)
	mp[n] = true
	for {
		n = next(n)
		if _, ok := mp[n]; ok {
			if n == 1 {
				return true
			}
			break
		} else {
			mp[n] = true
		}
	}
	return false
}

func next(n int) int {
	sum := 0
	for n >= 10 {
		remains := n % 10
		sum += remains * remains
		n = n / 10
	}
	return sum + n*n
}
