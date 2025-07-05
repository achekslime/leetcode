package medium

func CanJump(nums []int) bool {
	finish := len(nums) - 1
	position := finish

	for i := position - 1; i >= 0 && position > 0; i-- {
		if i+nums[i] >= position {
			position = i
		}
	}

	if position == 0 {
		return true
	}
	return false
}
