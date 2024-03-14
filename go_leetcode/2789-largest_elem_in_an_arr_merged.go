package go_leetcode

func maxArrayValue(nums []int) int64 {
	maxValue := nums[len(nums)-1]
	for i := len(nums) - 1; i > 0; i-- {
		if nums[i] >= nums[i-1] {
			nums[i-1] += nums[i]
		}
		if nums[i-1] > maxValue {
			maxValue = nums[i-1]
		}
	}
	return int64(maxValue)
}
