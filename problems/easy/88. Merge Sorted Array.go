package easy

func merge(nums1 []int, m int, nums2 []int, n int) {
	reserve := make([]int, len(nums1))
	copy(reserve, nums1)
	j, k := 0, 0
	for i := 0; i < m+n; i++ {
		if j >= m {
			nums1[i] = nums2[k]
			k++
			continue
		}
		if k >= n {
			nums1[i] = reserve[j]
			j++
			continue
		}
		if reserve[j] < nums2[k] {
			nums1[i] = reserve[j]
			j++
		} else {
			nums1[i] = nums2[k]
			k++
		}
	}
}
