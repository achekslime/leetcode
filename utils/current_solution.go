package utils

func ClimbStairs(n int) int {
	d := make([]int, n+1)
	d[0] = 1
	d[1] = 1
	for i := 2; i <= n; i++ {
		d[i] = d[i-1] + d[i-2]
	}
	return d[n]
}
