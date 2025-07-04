package easy

func removeDuplicates(nums []int) int {
	if len(nums) <= 1 {
		return 1
	}

	mp := make(map[int]struct{}, len(nums))
	next := 0
	ans := 0
	for i := 0; next < len(nums); i++ {
		if uniq(mp, nums[i]) {
			ans++
			if i == next {
				next++
			}
			continue
		}
		swap(nums, i, next)
		if !uniq(mp, nums[i]) {
			i--
		} else {
			ans++
		}
		next++
	}
	return ans
}

func uniq(mp map[int]struct{}, val int) bool {
	if _, ok := mp[val]; ok {
		return false
	}
	mp[val] = struct{}{}
	return true
}
