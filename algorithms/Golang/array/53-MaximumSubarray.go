package array

import "math"

// O(n^2)
func maxSubArray(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	length, max, temp := len(nums), nums[0], 0
	for i := 0; i < length; i++ {
		for j := i + 1; j <= length; j++ {
			temp = temp + nums[j-1]
			if temp > max {
				max = temp
			}
		}
		temp = 0
	}
	return max
}

// solution2
// URL https://en.wikipedia.org/wiki/Maximum_subarray_problem
// O(n)
func maxSubArray2(nums []int) int {
	max_ending_here, max_so_far := nums[0], nums[0]
	length := len(nums)
	for i := 1; i < length; i++ {
		max_ending_here = int(math.Max(float64(nums[i]), float64(max_ending_here+nums[i])))
		max_so_far = int(math.Max(float64(max_so_far), float64(max_ending_here)))
	}
	return max_so_far
}
