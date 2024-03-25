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
	maxHere, maxSoFar := 0, math.MinInt
	length := len(nums)
	for i := 0; i < length; i++ {
		maxHere = calMax(nums[i], maxHere+nums[i])
		maxSoFar = calMax(maxHere, maxSoFar)
	}
	return maxSoFar
}

func calMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}
