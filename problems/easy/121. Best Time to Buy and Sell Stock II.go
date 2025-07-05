package easy

func MaxProfit2(prices []int) int {
	n := len(prices)
	lowest := prices[0]
	highest := prices[0]
	ans := 0
	local := 0
	for i := 1; i < n; i++ {
		local = max(local, prices[i]-lowest)
		if prices[i] > highest {
			highest = prices[i]
		}
		if prices[i] < lowest || prices[i] < highest {
			if local > 0 {
				ans += local
			}
			lowest = prices[i]
			highest = prices[i]
			local = 0
		}
	}
	ans = max(ans, ans+local)
	return ans
}
