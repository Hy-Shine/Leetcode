package go_leetcode

func moveZeroes(nums []int) {
	slow := 0
	fast := 0
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
