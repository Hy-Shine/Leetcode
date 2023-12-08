package go_leetcode

// problems 200-300 solutions

// No.268
func missingNumber(nums []int) int {
	n := len(nums)
	var m = make(map[int]struct{}, len(nums))
	for _, v := range nums {
		m[v] = struct{}{}
	}

	for i := 0; i <= n; i++ {
		if _, ok := m[i]; !ok {
			return i
		}
	}
	return -1
}

// No.283
func moveZeroes(nums []int) {
	slow, fast := 0, 0
	for fast < len(nums) {
		if nums[fast] == 0 {
			fast++
			continue
		}
		if nums[slow] == 0 && nums[fast] != 0 {
			nums[slow], nums[fast] = nums[fast], nums[slow]
		}
		slow++
		fast++
	}
}
