package medium

var mp map[int]int8

func removeDuplicates2(nums []int) int {
	if len(nums) <= 1 {
		return 1
	}
	initmap(len(nums))
	next := 0
	ans := 0
	for i := 0; next < len(nums); i++ {
		if uniq2(nums[i]) {
			ans++
			if i == next {
				next++
			}
			continue
		}
		swap(nums, i, next)
		if !uniq2(nums[i]) {
			i--
		} else {
			ans++
		}
		next++
	}
	return ans
}

func initmap(lenMp int) {
	mp = make(map[int]int8, lenMp)
}

func uniq2(val int) bool {
	v, ok := mp[val]
	if !ok || (ok && v < 2) {
		mp[val]++
		return true
	}
	return false
}

func swap(nums []int, i, r int) {
	if i == r {
		return
	}
	nums[i] += nums[r]
	nums[r] = nums[i] - nums[r]
	nums[i] = nums[i] - nums[r]
}
