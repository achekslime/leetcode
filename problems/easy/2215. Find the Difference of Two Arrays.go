package easy

func findDifference(nums1 []int, nums2 []int) [][]int {
	set1 := make(map[int]struct{}, len(nums1))
	set2 := make(map[int]struct{}, len(nums2))

	for _, v := range nums1 {
		set1[v] = struct{}{}
	}
	for _, v := range nums2 {
		set2[v] = struct{}{}
	}

	ans := make([][]int, 2)
	for key, _ := range set1 {
		if _, ok := set2[key]; !ok {
			ans[0] = append(ans[0], key)
		}
	}
	for key, _ := range set2 {
		if _, ok := set1[key]; !ok {
			ans[1] = append(ans[1], key)
		}
	}
	return ans
}
