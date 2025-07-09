package easy

func findDifference(nums1 []int, nums2 []int) [][]int {
	n, m := len(nums1), len(nums2)

	mp1 := make(map[int]struct{}, n)
	for _, v := range nums1 {
		mp1[v] = struct{}{}
	}
	mp2 := make(map[int]struct{}, m)
	for _, v := range nums2 {
		mp2[v] = struct{}{}
	}

	mpAns1 := make(map[int]struct{})
	mpAns2 := make(map[int]struct{})
	ans := make([][]int, 2)
	for _, v := range nums1 {
		if _, ok := mp2[v]; ok {
			continue
		}
		if _, ok := mpAns1[v]; ok {
			continue
		}
		ans[0] = append(ans[0], v)
		mpAns1[v] = struct{}{}
	}
	for _, v := range nums2 {
		if _, ok := mp1[v]; ok {
			continue
		}
		if _, ok := mpAns2[v]; ok {
			continue
		}
		ans[1] = append(ans[1], v)
		mpAns2[v] = struct{}{}
	}

	return ans
}
