package easy

func merge(nums1 []int, m int, nums2 []int, n int) {
	j, k := m-1, n-1
	i := m + n - 1
	for ; j >= 0 && k >= 0; i-- {
		if nums2[k] >= nums1[j] {
			nums1[i] = nums2[k]
			k--
		} else {
			nums1[i] = nums1[j]
			j--
		}
	}
	if k >= 0 {
		copy(nums1[:i+1], nums2[:k+1])
	}
}
