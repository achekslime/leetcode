package solution

func MajorityElement(nums []int) int {
	frequency := make(map[int]int)
	for i := range nums {
		frequency[nums[i]]++
	}

	maxFreq := frequency[nums[0]]
	maxValue := nums[0]
	for i := range frequency {
		if frequency[i] > maxFreq {
			maxFreq = frequency[i]
			maxValue = i
		}
	}

	return maxValue
}
