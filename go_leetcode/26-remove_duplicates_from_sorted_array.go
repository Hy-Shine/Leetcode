package go_leetcode

func removeDuplicates(nums []int) int {
	if len(nums) <= 1 {
		return len(nums)
	}

	slow, fast := 0, 1
	for slow <= fast && fast < len(nums) {
		if nums[slow] == nums[fast] {
			fast++
		} else {
			slow++
			nums[slow] = nums[fast]
			fast++
		}
	}
	return slow + 1
}
