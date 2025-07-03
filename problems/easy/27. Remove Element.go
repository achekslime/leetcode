package easy

func removeElement(nums []int, val int) int {
	last := len(nums) - 1
	ans := 0
	for i := 0; i <= last; i++ {
		if nums[i] == val {
			swap(nums, i, last)
			ans++
			last--
			i--
		}
	}
	return len(nums) - ans
}

func swap(nums []int, i, r int) {
	if i == r {
		return
	}
	nums[i] += nums[r]
	nums[r] = nums[i] - nums[r]
	nums[i] = nums[i] - nums[r]
}
