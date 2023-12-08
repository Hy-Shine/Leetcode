package go_leetcode

func sortedSquares(nums []int) []int {
	if len(nums) == 1 {
		return []int{nums[0] * nums[0]}
	}

	for i := range nums {
		nums[i] = nums[i] * nums[i]
	}

	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums)-i-1; j++ {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
		}
	}

	return nums
}
