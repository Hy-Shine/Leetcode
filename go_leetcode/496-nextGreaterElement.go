package go_leetcode

// time require: o(m+n)

func NextGreaterElement(nums1 []int, nums2 []int) []int {
	ans := make([]int, 0, len(nums1))
	for _, v := range nums1 {
		ans = append(ans, nextElement(v, nums2))
	}
	return ans
}

func nextElement(target int, list []int) int {
	exists := false
	for i, v := range list {
		if v == target {
			exists = true
			continue
		}
		if exists {
			if v > target && i <= len(list)-1 {
				return list[i]
			}
		}
	}
	return -1
}
