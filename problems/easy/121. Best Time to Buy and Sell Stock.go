package easy

func MaxProfit(prices []int) int {
	n := len(prices)
	min, max := prices[0], prices[0]
	ans := 0
	for i := 1; i < n; i++ {
		if prices[i] > max {
			max = prices[i]
			if max-min > ans {
				ans = max - min
			}
		}
		if prices[i] < min {
			min = prices[i]
			max = prices[i]
		}
	}
	return ans
}
