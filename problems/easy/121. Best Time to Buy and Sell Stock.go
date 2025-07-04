package easy

func MaxProfit(prices []int) int {
	n := len(prices)
	lowest := prices[0]
	ans := 0
	for i := 1; i < n; i++ {
		ans = max(ans, prices[i]-lowest)
		if prices[i] < lowest {
			lowest = prices[i]
		}
	}
	return ans
}
