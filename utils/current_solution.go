package utils

func MissingNumber(nums []int) int {
	n := len(nums)
	sum := 0
	sumActual := 0
	for i, v := range nums {
		sum += i
		sumActual += v
	}
	sum += n
	return sum - sumActual
}
