package medium

func jump(nums []int) int {
	n := len(nums)

	// динамика
	d := make([]int, n)
	d[n-1] = 0

	for i := n - 2; i >= 0; i-- {
		j := findMin(d, i, i+nums[i])
		if i == j {
			d[i] = 1000000
		} else {
			d[i] = d[j] + 1
		}
	}
	return d[0]
}

func findMin(d []int, left int, right int) int {
	shortest := 1000000
	shortesti := left
	for i := left + 1; i <= right && i < len(d); i++ {
		if d[i] < shortest {
			shortest = d[i]
			shortesti = i
		}
	}
	return shortesti
}
